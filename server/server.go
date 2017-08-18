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
	//r.HandleFunc("/env/{id}", getEnv).Methods("GET")
	//r.HandleFunc("/env/{id}", updateEnv).Methods("PUT")
	//r.HandleFunc("/env/{status}", setEnvStatus).Methods("POST")

    log.Fatal(http.ListenAndServe(":4321", r))
}

