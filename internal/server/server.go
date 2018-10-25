package server

import (
	"net"
	"os"
	"os/signal"

	"github.com/dsbyrne/conjurd/internal/conjur"
	"github.com/rs/zerolog/log"
)

const (
	socketPath = "/var/run/conjurd.sock"
)

func Listen() error {
	socket, err := net.ListenUnix("unix", &net.UnixAddr{Name: socketPath})
	if err != nil {
		return err
	}

	conjur.Initialize()

	signalReceiver := make(chan os.Signal, 1)
	signal.Notify(signalReceiver, os.Interrupt)
	go func() {
		signal := <-signalReceiver
		log.Info().Msgf("shutting down after receiving signal %v", signal)

		err := socket.Close()
		if err != nil {
			log.Debug().
				Err(err).
				Msg("failed to close socket")
		}

		os.Exit(0)
	}()

	for {
		conn, err := socket.AcceptUnix()
		if err != nil {
			if _, ok := err.(*net.OpError); ok {
				log.Info().
					Msg("listener shutting down due to closed connection")
				return nil
			}

			log.Error().
				Err(err).
				Msgf("failed to accept connection, %T", err)
			continue
		}
		conn.Write(conjur.GetToken())
		conn.Close()
	}
}
