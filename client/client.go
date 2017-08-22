package client

import (
    "os"
    "net/http"
    "path"
    "time"
    "encoding/json"
    "bytes"
    "fmt"
    "bufio"
    _"io/ioutil"
    "log"

    "github.com/cheggaaa/pb"
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
                    Timeout: time.Second * 10000,
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
    resp, err := c.Do(req)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    defer resp.Body.Close()

    statusMap := make(map[string]interface{})
    layersProg := make(map[string]*pb.ProgressBar)
    uptodate := false

    fmt.Println("Pulling image from registry...")
    pool, err := pb.StartPool()
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        json.Unmarshal(scanner.Bytes(), &statusMap)
        id := statusMap["layer"].(string)
        progress := statusMap["progress"].(float64)
        if id == "none" {
            uptodate = true
            break
        }
        if _, ok := layersProg[id]; ok {
            layersProg[id].Set(int(progress*100))
        } else {
            layersProg[id] = pb.New(100).Prefix(id+": ")
            layersProg[id].ShowTimeLeft = false
            layersProg[id].ShowCounters = false
            pool.Add(layersProg[id])
        }
    }
    pool.Stop()
    if uptodate {
        fmt.Printf("\r%s          \n", "Image is up to date!")
    }
    fmt.Printf("Environment created! To start, run \n  $ repscie env start [envname]\n")
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
