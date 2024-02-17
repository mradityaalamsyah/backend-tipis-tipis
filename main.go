package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter,r *http.Request) {
    w.Write([]byte("Tipis-Tipis captt"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Println("Starting web server at :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
