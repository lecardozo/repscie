package main

import (
    _"fmt"
    "log"
    "net/http"
    "encoding/json"
    _"io/ioutil"
    "io"
    "bufio"
    _"os"
    _"strings"
    "github.com/fsouza/go-dockerclient"

    "github.com/lecardozo/repsci/api/environment"
)

func createEnv(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var env environment.Environment
    err := decoder.Decode(&env)
    if err != nil {
        log.Fatal(err)
    }

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
    if err != nil {
		panic(err)
    }

    read, write := io.Pipe()
    opts := docker.PullImageOptions{
        Repository: "postgres",
        Tag: "latest",
        OutputStream: write,
        RawJSONStream: true,
    }

    // start goroutine pulling image and writing to write
    go func() {
        defer write.Close()
        err := client.PullImage(opts, docker.AuthConfiguration{})
        if err != nil {
            log.Fatal(err)
        }
    }()

    scanner := bufio.NewScanner(read)
    for scanner.Scan() {
        w.Write([]byte(scanner.Text()))
        w.(http.Flusher).Flush()
    }
    return 
}
