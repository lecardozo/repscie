package client

import (
    "net/http"
    "github.com/lecardozo/repsci/helper"
    "strconv"
)

const DefaultBaseURL = "http://localhost"
const DefaultPort = 4321

// Version is the current library's version: sent with User-Agent
const Version = "0.1"

// Client can create smite session objects and interact with the smite API
type RSClient struct {
    Host string
    Port int
    http.Client
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

func (c RSClient) CreateEnv() (string, error) {
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
