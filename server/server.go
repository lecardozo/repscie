package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/project", initProject).Methods("POST")
	r.HandleFunc("/project/{id}", getProject).Methods("GET")

    r.HandleFunc("/env", createEnv).Methods("POST")
	r.HandleFunc("/env/{id}", updateEnv).Methods("POST")
    //r.HandleFunc("/env/{id}", getEnv).Methods("GET")

    log.Fatal(http.ListenAndServe(":4321", r))
}

