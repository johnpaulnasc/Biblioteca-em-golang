package models

import (
    "github.com/johnpaulnasc/biblioteca-golang/internal/database"
)

type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Genre       string `json:"genre"`
    PublicationYear int `json:"publication_year"`
    Author      string `json:"author"`
    Publisher   string `json:"publisher"`
}

func CreateBook(b *Book) error {
    db := database.GetDB()
    stmt, err := db.Prepare("INSERT INTO books(title, genre, publication_year, author, publisher) VALUES(?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(b.Title, b.Genre, b.PublicationYear, b.Author, b.Publisher)
    if err != nil {
        return err
    }
    return nil
}

func UpdateBook(b *Book) error {
    db := database.GetDB()
    stmt, err := db.Prepare("UPDATE books SET title = ?, genre = ?, publication_year = ?, author = ?, publisher = ? WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(b.Title, b.Genre, b.PublicationYear, b.Author, b.Publisher, b.ID)
    if err != nil {
        return err
    }
    return nil
}

func DeleteBook(id int) error {
    db := database.GetDB()
    stmt, err := db.Prepare("DELETE FROM books WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
