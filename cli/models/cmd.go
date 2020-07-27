package models

import (
    "fmt"
    "os/exec"
)

func RunCommand(c string, a []string) []byte {

    out, err := exec.Command(c, a...).CombinedOutput()
    if err != nil {
        fmt.Println( "Error executing command:", err )
    }

    return out
}
