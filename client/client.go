package client

import (
    "os"
    "net/http"
    "path"
    "log"
    "time"
    "encoding/json"
    "bytes"
    "fmt"

    "github.com/lecardozo/repsci/helper"
    "github.com/lecardozo/repsci/api/project"
)

const DefaultBaseURL = "http://localhost"
const DefaultPort = 4321

// Version is the current library's version: sent with User-Agent
const Version = "0.1"

// Client can create smite session objects and interact with the smite API
type RSClient struct {
    host string
    *http.Client
}

func NewRSClient(host string) (*RSClient, error) {
    if host == "" {
        host = "http://localhost:4321"
    }

    return &RSClient{
                host,
                &http.Client{
                    Timeout: time.Second * 10,
                },
    }, nil
}

func (c RSClient) GetProjects() (string, error) {
    const path = "/project"
    const method = "GET"
    url := helper.Cat(c.host, path)
    //resp, err := c.Get(url)
    return url, nil
}

// This functions verifies if project already exists and, if not,
// initializes.
func (c *RSClient) InitProject(config *project.Config) (string, error) {
    exists, err := project.ProjectExists(config.Location)
    if exists {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return "", err
    }

    os.MkdirAll(path.Join(config.Location, ".repscie"), os.ModePerm)
    fmt.Printf("Initializing project at %s\n",config.Location)

    // POST new project
    const path = "/project"
    const method = "POST"
    url := helper.Cat(c.host, path)

    js, err := json.Marshal(config)
    req , err := http.NewRequest(method, url, bytes.NewBuffer(js))
    req.Header.Set("Content-Type", "application/json")
    resp, err := c.Do(req)
    if err != nil {
        return "", err
    }

    // Decode server response
    proj := project.Project{}
    err = json.NewDecoder(resp.Body).Decode(&proj)
    if err != nil {
        return "", err
    }

    return proj.ID, nil
}



func (c RSClient) ListEnvs() (string, error) {
    const path = "/envs"
    const method = "GET"
    url := helper.Cat(c.host, path)
    //resp, err := c.Get(url)
    return url, nil
}


func (c RSClient) CreateEnv(configfile string) (string, error) {
    if _, err := os.Stat(configfile); os.IsNotExist(err) {
        log.Fatalf("File %s does not exist", configfile)
    }

    const path = "/env"
    const method = "POST"
    url := helper.Cat(c.host, path)

    //resp, err := c.Get(url)
    return url, nil
}

func (c RSClient) UpdateEnv(id string) (string, error) {
    const path = "/env"
    const method = "PUT"
    url := helper.Cat(c.host, path)
    //resp, err := c.Get(url)
    return url, nil
}
