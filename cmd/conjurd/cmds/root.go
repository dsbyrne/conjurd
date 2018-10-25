// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmds

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "conjurd",
	Short: "Conjur machine identity",
	Long: `A daemon which serves as machine identity for Conjur. Access tokens are
available to clients over the unix domain socket located at /var/run/conjurd.sock.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		zerolog.TimeFieldFormat = ""
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if Debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		if Pretty {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}
	},
}

// Debug enabled debug logging
var Debug bool

// Pretty enables pretty output
var Pretty bool

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug output")
	rootCmd.PersistentFlags().BoolVarP(&Pretty, "pretty", "p", false, "pretty output")
}
