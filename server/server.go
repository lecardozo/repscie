package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func newProject(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New project created\n",)
    w.WriteHeader(http.StatusOK)
}

func getProject(w http.ResponseWriter, r *http.Request) {
    // retrieve project data
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Project: %v\n", vars["id"])
    w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/project", newProject).Methods("POST")
	r.HandleFunc("/project/{id}", getProject).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", r))
}
