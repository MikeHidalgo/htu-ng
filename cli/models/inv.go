package models

import (
    "fmt"
    "strconv"

    "htu-ng/cli/db"
)

var commands [][]string

func SetCommands() [][]string {

    cpu := []string{"dmidecode", "-t", "processor"}
    ram := []string{"dmidecode", "-t", "17"}
    sys := []string{"dmidecode", "-t", "system"}
    cha := []string{"dmidecode", "-t", "chassis"}
    bio := []string{"dmidecode", "-t", "bios"}
    pwr := []string{"dmidecode", "-t", "39"}

    commands = append(commands, cpu, ram, sys, cha, bio, pwr)

    return commands
}

func INV() {

    database, _ := db.Init()
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, output TEXT)")
    statement.Exec()

    statement, _ = database.Prepare("INSERT INTO commands (command, output) VALUES (?, ?)")

    for _, cmd := range SetCommands() {
        statement.Exec(cmd[2], RunCommand(cmd[0], cmd[1:]))
    }

    rows, _ := database.Query("SELECT id, command, output FROM commands")
    var id int
    var command string
    var output string
    for rows.Next() {
        rows.Scan(&id, &command, &output)
        fmt.Println("")
        fmt.Println("# .............................")
        fmt.Println(strconv.Itoa(id) + ": " + command)
        fmt.Println(output)
    }
}
