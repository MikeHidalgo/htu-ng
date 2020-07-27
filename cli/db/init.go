package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)
var path = "./htu.db"

func Init() (*sql.DB, error) {
    return sql.Open("sqlite3", path)
}
