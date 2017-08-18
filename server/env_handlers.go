package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    _"io/ioutil"
    "io"
    "bufio"
    "strings"
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
        Repository: "rocker/r-ver",
        Tag: "latest",
        OutputStream: write,
        RawJSONStream: false,
    }

    go client.PullImage(opts, docker.AuthConfiguration{})
    go func(reader io.Reader) {
        scanner := bufio.NewScanner(reader)
        for scanner.Scan() {
            fmt.Printf("%s \n", scanner.Text())
            if strings.Contains(scanner.Text(), "up to date") {
                break
            }
        }
        fmt.Printf("%s \n", "Acabooou")

    }(read)
    //imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
    //for _, img := range imgs {
    //}
//    env.Status = "CREATED"
//
//    js, err := json.Marshal(proj)
//    if err != nil {
//        http.Error(w, err.Error(), http.StatusInternalServerError)
//        return
//    }
//
//    w.Header().Set("Content-Type", "application/json")
//    w.WriteHeader(http.StatusOK)
//    w.Write(js)
}
