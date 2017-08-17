package client

import (
    "os"
    "net/http"
    "encoding/json"
    "github.com/lecardozo/repsci/api/environment"
    "github.com/lecardozo/repsci/api/project"
    "strconv"
)

const DefaultBaseURL = "http://localhost"
const DefaultPort = 4321

// Version is the current library's version: sent with User-Agent
const Version = "0.1"

// Client can create smite session objects and interact with the smite API
type RSClient struct {
    host string
    client *http.Client
}

func NewRSClient(host string) (*RSClient, error) {
    if host == nil {
        host = "http://localhost:4321"
    }

    return &RSClient{
                host: host,
                client: &http.Client{
                    Timeout: time.Second * 10
                }

    }, nil
}

func (c RSClient) GetProjects() (string, error) {
    const path = "/project"
    const method = "GET"
    url := helper.Cat(c.Host, ":", strconv.Itoa(c.Port), path)
    //resp, err := c.Get(url)
    return url, nil
}

func (c RSClient) InitProject() (string, error) {
    const path = "/project"
    const method = "POST"
    url := helper.Cat(c.Host, ":", strconv.Itoa(c.Port), path)
    //resp, err := c.Get(url)
    return url, nil
}

func (c RSClient) ListEnvs() (string, error) {
    const path = "/envs"
    const method = "GET"
    url := helper.Cat(c.Host, ":", strconv.Itoa(c.Port), path)
    //resp, err := c.Get(url)
    return url, nil
}


func (c RSClient) CreateEnv(configfile string) (string, error) {
    if _, err := os.Stat(configfile); os.IsNotExist(err) {
        log.Fatalf("File %s does not exist", configfile)
    }

    envconf := environment.ConfigFromFile(configfile)
    const path = "/env"
    const method = "POST"
    url := helper.Cat(c.Host, ":", strconv.Itoa(c.Port), path)

    //resp, err := c.Get(url)
    return url, nil
}

func (c RSClient) UpdateEnv(id string) (string, error) {
    const path = "/env"
    const method = "PUT"
    url := helper.Cat(c.Host, ":", strconv.Itoa(c.Port), path)
    //resp, err := c.Get(url)
    return url, nil
}
