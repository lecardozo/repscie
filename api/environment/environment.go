package environment

import (
    "strings"
)

type Environment struct {
    Lang string // python or r
    Version string
    ImageID string
    BaseImage string // image name
    Status string // created, running, stopped
}


func DefaultEnv(lang string) *Environment {
    lang = strings.ToLower(lang)
    var image string
    switch lang {
        case "r":
            image = "rocker/r-ver"
        case "python" :
            image = "jupyter/minimal-notebook"
    }
    return &Environment{lang,
                        "latest",
                        "",
                        image,
                        ""}
}
