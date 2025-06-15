package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSource string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dataSource)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec("PRAGMA foreign_keys = ON;")
    if err != nil {
        return nil, err
    }

    if err := migrate(db); err != nil {
        return nil, err
    }

    return db, nil
}

func migrate(db *sql.DB) error {
    userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL
    );
    `

    fileTable := `
    CREATE TABLE IF NOT EXISTS files (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        size INTEGER NOT NULL,
        uploaded TIMESTAMP NOT NULL,
        url TEXT NOT NULL,
        user_id INTEGER NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );
    `

    if _, err := db.Exec(userTable); err != nil {
        return err
    }

    if _, err := db.Exec(fileTable); err != nil {
        return err
    }

    log.Println("SQLite migrations applied successfully.")
    return nil
}