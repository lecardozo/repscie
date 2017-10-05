package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/satori/go.uuid"
    "github.com/lecardozo/repsci/api/project"
)

func initProject(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var proj project.Project
    err := decoder.Decode(&proj)
    if err != nil {
        log.Fatal(err)
    }

    id := uuid.NewV4()

    proj.ID = id.String()
    proj.Status = "CREATED"

    js, err := json.Marshal(proj)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(js)
}

func getProject(w http.ResponseWriter, r *http.Request) {
    // retrieve project data
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Project: %v\n", vars["id"])
    w.WriteHeader(http.StatusOK)
}
