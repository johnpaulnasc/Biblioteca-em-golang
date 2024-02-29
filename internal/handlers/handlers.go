package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/johnpaulnasc/biblioteca-golang/internal/database"
    "github.com/johnpaulnasc/biblioteca-golang/internal/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    db := database.GetDB()
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var b models.Book
        if err := rows.Scan(&b.ID, &b.Title, &b.Genre, &b.PublicationYear, &b.Author, &b.Publisher); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        books = append(books, b)
    }

    json.NewEncoder(w).Encode(books)
}
