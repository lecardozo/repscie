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
    "os"
    "path"
    "path/filepath"
    "log"
    "github.com/spf13/cobra"
    "github.com/lecardozo/repsci/client"
)

var ProjectCmd = &cobra.Command{
    Use:   "project",
    Short: "Repsci project management",
    Long: `This subcommand is used to manage repsci projects`,
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize repsci project",
    Long: `Initializes your repsci project structure`,
    Run: initProject,
}

var listCmd = &cobra.Command{
    Use:   "ls",
    Short: "List repsci projects",
    Long: `Lists all repsci projects on current machine`,
    Run: listProjects,
}

func initProject(cmd *cobra.Command, args []string) {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    if args[0] == "." {
        os.Mkdir(path.Join(dir, "data"), os.ModePerm)
        fmt.Printf("Initializing project at %s\n", dir)
    } else if path.IsAbs(args[0]) {
        os.MkdirAll(path.Join(args[0], "data"), os.ModePerm)
        fmt.Printf("Initializing project at %s\n", args[0])
    } else {
        os.MkdirAll(path.Join(dir, args[0], "data"), os.ModePerm)
        fmt.Printf("Initializing project at %s\n",
                           path.Join(dir, args[0]))
    }
}

func listProjects(cmd *cobra.Command, args []string) {

    client := client.RSClient{
        Host: "localhost",
        Port: 4321,
    }

    resp, err := client.GetProjects()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}

func init() {
    ProjectCmd.AddCommand(initCmd)
    ProjectCmd.AddCommand(listCmd)
    RootCmd.AddCommand(ProjectCmd)
}
