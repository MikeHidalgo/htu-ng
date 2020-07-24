package models

import (
    "fmt"
    "os/exec"
    "strconv"

    "htu-ng/cli/db"
    "htu-ng/cli/models"
)


type Commands struct {
    CPU  []string
    RAM  []string
    SYS  []string
    CHA  []string
    BIO  []string
    PWR  []string
}

func SetCommands() *Commands {
    return &Command{
        CPU: {"dmidecode", "-t", "processor"},
        RAM: {"dmidecode", "-t", "17"},
        SYS: {"dmidecode", "-t", "system"},
        CHA: {"dmidecode", "-t", "chassis"},
        BIO: {"dmidecode", "-t", "bios"},
        PWR: {"dmidecode", "-t", "39"},
    }
}

func RunCommand(c string, a []string) []byte {

    out, err := exec.Command(c, a...).CombinedOutput()
    if err != nil {
        fmt.Println( "Error:", err )
    }

    // fmt.Printf("%s\n", string(out))
    return out
}

func INV() {

    database, _ := db.DB("./htu.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, output TEXT)")
    statement.Exec()
    statement, _ = database.Prepare("INSERT INTO commands (command, output) VALUES (?, ?)")

    for _, cmd := range commands {
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
