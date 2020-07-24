package models

import (
    "htu-ng/cli/db"
)

func Summary() {

    init := "CREATE TABLE IF NOT EXISTS system (id INTEGER, component TEXT PRIMARY KEY, output TEXT)"

    database, _ := db.DB("./htu.db")
    statement, _ := database.Prepare(init)
    statement.Exec()

    statement, _ = database.Prepare("INSERT INTO system (component, output) VALUES (?, ?)")
    statement.Exec("System Information", RunCommand("dmidecode", []string{"-t", "1"}))
}
