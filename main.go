package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist/")))

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
