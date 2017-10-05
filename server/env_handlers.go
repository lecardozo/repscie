package main

import (
    "fmt"
    "log"
    "net/http"
    "context"
    "encoding/json"
    _"io/ioutil"
    _"io"
    "bufio"
    _"os"
    "strings"
    _"strconv"
    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
    "github.com/moby/moby/api/types/container"


    "github.com/lecardozo/repsci/api/environment"
)

func createEnv(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var env environment.Environment
    err := decoder.Decode(&env)
    if err != nil {
        log.Fatal(err)
    }

	cli, err := client.NewEnvClient()
    if err != nil {
		panic(err)
    }

    opts := types.ImagePullOptions{}
    readc, err := cli.ImagePull(context.Background(), env.BaseImage, opts)
    if err != nil {
        panic(err)
    }
    defer readc.Close()

    flusher, ok := w.(http.Flusher)
    if !ok {
      fmt.Println("expected http.ResponseWriter to be an http.Flusher")
    }

    statusMap := make(map[string]interface{})
    layerProg := make(map[string]float64)

    scanner := bufio.NewScanner(readc)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        json.Unmarshal(scanner.Bytes(), &statusMap)
        status := statusMap["status"].(string)
        switch {
            case strings.Contains(status,"Image is up to date"):
                fmt.Fprintf(w, "{\"layer\": \"%s\", \"progress\": %f}\n",
                                    "none", 0.0)
            case status == "Already exists":
            case status == "Pulling fs layer":
                fmt.Fprintf(w, "{\"layer\": \"%s\", \"progress\": %f}\n",
                                statusMap["id"].(string), 0.0)
                layerProg[statusMap["id"].(string)] = 0.0
            case status == "Waiting":
                layerProg[statusMap["id"].(string)] = 0.0
            case status == "Downloading":
                progDet := statusMap["progressDetail"].(map[string]interface{})
                cur := progDet["current"].(float64)
                tot := progDet["total"].(float64)
                layerProg[statusMap["id"].(string)] = cur
                fmt.Fprintf(w, "{\"layer\": \"%s\", \"progress\": %f}\n",
                             statusMap["id"].(string), cur/tot)

            case status == "Verifying Checksum":
            case status == "Download complete":
            case status == "Extracting":
            case status == "Pull complete":
        }
        flusher.Flush()
    }
    fmt.Fprint(w, )
}

func updateEnv(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var env environment.Environment
    err := decoder.Decode(&env)
    if err != nil {
        log.Fatal(err)
    }

	cli, err := client.NewEnvClient()
    if err != nil {
		panic(err)
    }

    cconf := container.Config{
                 Hostname: "repscie",
                 User: "root",
                 AttachStdin: true,
                 AttachStdout: true,
                 AttachStderr: true,
                 Tty: true,
                 OpenStdin: true,
                 Cmd:
                 ArgsEscaped     bool
                 Image           string
                 Volumes         map[string]struct{}
                 WorkingDir      string
                 Entrypoint      strslice.StrSlice
                 Shell           strslice.StrSlice
    }

    cid, err := cli.ContainerCreate(context.Background(),
                                    cconf,
                                    hostConfig *container.HostConfig,
                                    networkingConfig *network.NetworkingConfig,
                                    containerName string)
    if err != nil {
        panic(err)
    }
}

