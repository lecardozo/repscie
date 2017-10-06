package repscie

import (
    "os"
    "path"
    _"time"
    _"bytes"
    "fmt"
    "bufio"
    _"io/ioutil"
    _"log"
    "context"
    "encoding/json"
    "strings"

    "github.com/gosuri/uiprogress"
    "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"


    _"github.com/lecardozo/repscie/helper"
    "github.com/lecardozo/repscie/api/project"
)

// This functions verifies if project already exists and, if not,
// initializes.
func InitProject(config *project.Config) (error) {
    exists, err := project.ProjectExists(config.Location)
    if exists {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return err
    }

    os.MkdirAll(path.Join(config.Location, ".repscie"), os.ModePerm)
    fmt.Printf("Initializing project at %s\n",config.Location)

    // TODO: right config file

    return nil
}


func CreateEnv(lang string) (error) {
    dir, err := os.Getwd()
    exists, err := project.ProjectExists(dir)
    if !exists {
        fmt.Fprintf(os.Stderr, "Error: environment must be created inside project\n", )
        return err
    }

    cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
    opts := types.ImagePullOptions{}
    readc, err := cli.ImagePull(context.Background(), lang, opts)
    if err != nil {
        panic(err)
    }
    defer readc.Close()

    statusMap := make(map[string]interface{})
    layersProg := make(map[string]*uiprogress.Bar)
    uiprogress.Start()
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(readc)
    fmt.Println("Pulling Environment Image")
    for scanner.Scan() {
        json.Unmarshal(scanner.Bytes(), &statusMap)
        id := statusMap["id"].(string)
        status := statusMap["status"].(string)
        switch {
            case strings.Contains(status,"Image is up to date"):
            case status == "Already exists":
                layersProg[id].Set(50)
            case status == "Pulling fs layer":
                layersProg[id] = uiprogress.AddBar(100).AppendCompleted()
                layersProg[id].PrependFunc(func(b *uiprogress.Bar) string {
                    return fmt.Sprintf("%s: ", id)
                })
                layersProg[id].Width = 20
            case status == "Downloading":
                progDet := statusMap["progressDetail"].(map[string]interface{})
                cur := progDet["current"].(float64)
                tot := progDet["total"].(float64)
                layersProg[id].Set(int((cur/tot)*100))
            case status == "Verifying Checksum":
            case status == "Download complete":
            case status == "Extracting":
            case status == "Pull complete":
            default:
        }
    }
    fmt.Println("Environment ready to go!")
    return nil
}
