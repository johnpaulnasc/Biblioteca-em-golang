package main

import (
    "log"
    "net/http"

    "github.com/johnpaulnasc/biblioteca-golang/internal/handlers"
)

func main() {
    fs := http.FileServer(http.Dir("./web/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/books", handlers.GetBooks)

    log.Println("Iniciando servidor na porta 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
