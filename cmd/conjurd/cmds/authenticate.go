package cmds

import (
	"fmt"

	"github.com/dsbyrne/conjurd/pkg/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Retreive a Conjur access token",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := client.GetToken()
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to listen")
		}
		fmt.Println(string(token))
	},
}
