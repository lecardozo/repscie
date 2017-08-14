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

package cmd

import (
    "context"
	"fmt"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// playgroundCmd represents the playground command
var playgroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
}

var rstudioCmd = &cobra.Command{
	Use:   "rstudio",
	Short: "Starts R development environment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
    Run: startPlayground,
}

var ipythonCmd = &cobra.Command{
	Use:   "ipython",
	Short: "Starts ipython development environment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
    Run: startPlayground,
}


func startPlayground(cmd *cobra.Command, args []string) {

    var image string

    if cmd.Use == "ipython" {
        fmt.Println("Starting iPython environment")
        image = "ipython"
    } else if cmd.Use == "rstudio" {
        fmt.Println("Starting R environment")
        image = "rstudio"
    }

    fmt.Printf("Starting %s playground", image)

    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
    if err != nil {
        panic(err)
    }

    for _, image := range images {
        fmt.Printf("%s\n", image.Labels)
    }

}

func init() {

    RootCmd.AddCommand(playgroundCmd)
    playgroundCmd.AddCommand(rstudioCmd)
    playgroundCmd.AddCommand(ipythonCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playgroundCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playgroundCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
