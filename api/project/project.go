package project

import (
    "os"
    "errors"
    "path"
    "gopkg.in/yaml.v2"
    "github.com/lecardozo/repscie/helper"
    "github.com/lecardozo/repscie/api/environment"
    "log"
)

type Config struct {
    Name string
    Location string
    Backups Backup
}

type Backup struct {
    Freq string
    Provider string
    Status string
}

type Project struct {
    ID string
    Environment []environment.Environment
    Config Config
    Status string
}

func ProjectExists(dir string) (bool, error) {
    _, err := os.Stat(path.Join(dir, ".repscie"))
    if err != nil {
        if os.IsNotExist(err) {
            return false, nil
        } else {
            return true, err
        }
    } else {
        return true, errors.New("Project already exists")
    }
}

func ConfigFromFile(file string) *Config{
    config := helper.ReadConfigFile(file)
    newConf := Config{}
    err := yaml.Unmarshal([]byte(config), &newConf)
    if err != nil {
        log.Fatal("error")
    }
    return &newConf
}

func DefaultConfig(name string, dir string) *Config {
   return &Config{name, dir, Backup{"never", "none", "disabled"}}
}
