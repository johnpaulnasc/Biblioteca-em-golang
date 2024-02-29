package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("sqlite3", "./library.db")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        genre TEXT,
        publication_year INTEGER,
        author TEXT,
        publisher TEXT
    )`)
    if err != nil {
        log.Fatal(err)
    }
}

func GetDB() *sql.DB {
    return db
}
