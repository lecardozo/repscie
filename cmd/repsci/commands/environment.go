// Copyright © 2017 Lucas Cardozo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
    _"bufio"
    _"io"

    "github.com/lecardozo/repsci/client"
)

var EnvironmentCmd = &cobra.Command{
	Use:   "env",
	Short: "Command to manage environments",
	Long: "This command is used to create, update and destroy environments",
}

var createenvCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new environment",
	Long: "This command is used to create a new environment based on a rsenv.yml file",
    Run: createEnv,
}

var startenvCmd = &cobra.Command{
	Use:   "start",
	Short: "Creates new environment",
	Long: "This command is used to create a new environment based on a rsenv.yml file",
    Run: startEnv,
}

var updateenvCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates environment",
	Long: "This command is used to update your environment based on a rsenv.yml file",
    Run: updateEnv,
}


var host, configfile string

func init() {
    RootCmd.PersistentFlags().StringVarP(&host, "host", "H", "http://localhost:4321",
                                              "Daemon socket to connect to")
    createenvCmd.Flags().StringVarP(&configfile, "config", "c", "rsenv.yml",
                                                 `Environment configuration file
                                                 (default is ./rsenv.yml)`)
    updateenvCmd.Flags().StringVarP(&configfile, "config", "c", "rsenv.yml",
                                                 `Environment configuration file
                                                 (default is ./rsenv.yml)`)
    EnvironmentCmd.AddCommand(createenvCmd)
    EnvironmentCmd.AddCommand(startenvCmd)
    EnvironmentCmd.AddCommand(updateenvCmd)
    RootCmd.AddCommand(EnvironmentCmd)
}

func createEnv(cmd *cobra.Command, args []string) {
    rsclient, _ := client.NewRSClient(host)
    rsclient.CreateEnv(args[0])
}

func startEnv(cmd *cobra.Command, args []string) {
    fmt.Println("Environment is being started")
}

func updateEnv(cmd *cobra.Command, args []string) {
    fmt.Println("Environment has been updated")
}
