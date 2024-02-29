package models

type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Genre       string `json:"genre"`
    PublicationYear int `json:"publication_year"`
    Author      string `json:"author"`
    Publisher   string `json:"publisher"`
}