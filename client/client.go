package client

import (
    "os"
    "net/http"
    "path"
    "time"
    "encoding/json"
    "bytes"
    "fmt"

    "github.com/lecardozo/repsci/helper"
    "github.com/lecardozo/repsci/api/project"
    "github.com/lecardozo/repsci/api/environment"
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
func (c *RSClient) InitProject(config *project.Config) (*project.Project, error) {
    exists, err := project.ProjectExists(config.Location)
    if exists {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return nil, err
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
        return nil, err
    }

    // Decode server response
    proj := project.Project{}
    err = json.NewDecoder(resp.Body).Decode(&proj)
    if err != nil {
        return nil, err
    }

    return &proj, nil
}


func (c RSClient) CreateEnv(lang string) (*environment.Environment, error) {
    dir, err := os.Getwd()
    exists, err := project.ProjectExists(dir)
    if !exists {
        fmt.Fprintf(os.Stderr, "Error: environment must be created inside project\n", )
        return nil, err
    }

    const path = "/env"
    const method = "POST"
    url := helper.Cat(c.host, path)

    env := environment.DefaultEnv(lang)

    js, err := json.Marshal(env)
    req , err := http.NewRequest(method, url, bytes.NewBuffer(js))
    req.Header.Set("Content-Type", "application/json")
    resp, err := c.Do(req)
    if err != nil {
        return nil, err
    }

    // Decode server response
    err = json.NewDecoder(resp.Body).Decode(env)
    if err != nil {
        return nil, err
    }

    return env, nil
}

func (c RSClient) UpdateEnv(id string) (string, error) {
    const path = "/env"
    const method = "PUT"
    url := helper.Cat(c.host, path)
    //resp, err := c.Get(url)
    return url, nil
}

func (c RSClient) ListEnvs() (string, error) {
    const path = "/envs"
    const method = "GET"
    url := helper.Cat(c.host, path)
    //resp, err := c.Get(url)
    return url, nil
}
