package main

import (
    "net/http"
    "github/diploma/internal/handlers"
)

func main() {
    http.HandleFunc("/graph", handlers.GraphHandler)
    http.HandleFunc("/input", handlers.InputHandler)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
    http.ListenAndServe(":8000", nil)
}
