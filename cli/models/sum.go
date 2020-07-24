package models

import (
    "fmt"
    "strconv"

    "htu-ng/cli/db"
)

func Summary() {

    fmt.Println("summary called")

    init := "CREATE TABLE IF NOT EXISTS system (id INTEGER PRIMARY KEY, component TEXT, output TEXT)"

    database, _ := db.DB("./htu.db")
    statement, _ := database.Prepare(init)
    statement.Exec()

    statement, _ = database.Prepare("INSERT INTO system (component, output) VALUES (?, ?)")
    statement.Exec("System Information", RunCommand("dmidecode", []string{"-t", "1"}))

    rows, _ := database.Query("SELECT id, component, output FROM system")
    var id int
    var component string
    var output string
    for rows.Next() {
        rows.Scan(&id, &component, &output)
        fmt.Println("")
        fmt.Println("# .............................")
        fmt.Println(strconv.Itoa(id) + ": " + component)
        fmt.Println(output)
    }
}
