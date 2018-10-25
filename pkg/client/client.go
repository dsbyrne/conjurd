package client

import (
	"net"
)

// GetToken returns an access token from conjurd server
func GetToken() ([]byte, error) {
	conn, err := net.Dial("unix", "/var/run/conjurd.sock")
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}
