package db

import (
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

func DB() (*sql.DB, error) {

    database, err := sql.Open("sqlite3", "./htu.db")

    if err != nil {
        return nil, err
    }

    return database, nil
}
