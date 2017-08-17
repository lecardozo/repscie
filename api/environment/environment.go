package environment

import (
    "gopkg.in/yaml.v2"
    "log"
    "github.com/lecardozo/repsci/helper"
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

type Environment struct {
    Config Config
    ImageID string
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
