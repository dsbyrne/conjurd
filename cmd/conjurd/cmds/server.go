package cmds

import (
	"github.com/dsbyrne/conjurd/internal/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage the conjurd server",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Listen()
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to listen")
		}
	},
}
