package main

import (
    "log"
    "net/http"

    "github.com/johnpaulnasc/biblioteca-golang/internal/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
    r.HandleFunc("/books/create", handlers.CreateBookHandler).Methods("POST")
    r.HandleFunc("/books/{id}", handlers.UpdateBookHandler).Methods("PUT")
    r.HandleFunc("/books/{id}", handlers.DeleteBookHandler).Methods("DELETE")

    fs := http.FileServer(http.Dir("./web/static"))
    r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

    log.Println("Iniciando servidor na porta 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
