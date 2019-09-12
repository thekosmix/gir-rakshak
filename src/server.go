package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/users/{id}/locations/{from-date}/{to-date}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        fromDate := vars["from-date"]
        toDate := vars["to-date"]

        fmt.Fprintf(w, "You've requested the user's location: %s from-date %s and to-date %s\n", id, fromDate, toDate)
    })

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    fs := http.FileServer(http.Dir("static/"))
    r.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", r)
}