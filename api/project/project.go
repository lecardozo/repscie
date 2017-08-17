package project

import (
    "gopkg.in/yaml.v2"
    "github.com/lecardozo/repsci/helper"
    "github.com/lecardozo/repsci/api/environment"
    "log"
)

type Config struct {
    Deps struct {
        System struct {
            Pkgs []string
        }

        R struct {
            Pkgs []string
        }

        Python struct {
            Pkgs []string
        }
    }
}

type Project struct {
    Environment environment.Environment
    Config Config
    Status string
}


func ConfigFromFile(file string) Config{
    config := helper.ReadConfigFile(file)
    newConf := Config{}
    err := yaml.Unmarshal([]byte(config), &newConf)
    if err != nil {
            log.Fatal("error")
            }
            return newConf
}
