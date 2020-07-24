package models

import (
    "fmt"
    "os/exec"
    "strconv"

    "htu-ng/cli/db"
)

var commands [][]string

func SetCommands() *commands {

    commands = append(commands, {"dmidecode", "-t", "processor"})
    commands = append(commands, {"dmidecode", "-t", "17"})
    commands = append(commands, {"dmidecode", "-t", "system"})
    commands = append(commands, {"dmidecode", "-t", "chassis"})
    commands = append(commands, {"dmidecode", "-t", "bios"})
    commands = append(commands, {"dmidecode", "-t", "39"})

    return &commands
}

func RunCommand(c string, a []string) []byte {

    out, err := exec.Command(c, a...).CombinedOutput()
    if err != nil {
        fmt.Println( "Error:", err )
    }

    return out
}

func INV() {

    database, _ := db.DB("./htu.db")
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
