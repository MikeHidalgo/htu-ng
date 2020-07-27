package models

import (
    "fmt"
    "strconv"

    "htu-ng/cli/db"
)

func CodeTable() map[int]string {

    c := map[int]string{

         0: "BIOS",
         1: "System",
         2: "Base Board",
         3: "Chassis",
         4: "Processor",
         5: "Memory Controller",
         6: "Memory Module",
         7: "Cache",
         8: "Port Connector",
         9: "System Slots",
        10: "On Board Devices",
        11: "OEM Strings",
        12: "System Configuration Options",
        13: "BIOS Language",
        14: "Group Associations",
        15: "System Event Log",
        16: "Physical Memory Array",
        17: "Memory Device",
        18: "32-bit Memory Error",
        19: "Memory Array Mapped Address",
        20: "Memory Device Mapped Address",
        21: "Built-in Pointing Device",
        22: "Portable Battery",
        23: "System Reset",
        24: "Hardware Security",
        25: "System Power Controls",
        26: "Voltage Probe",
        27: "Cooling Device",
        28: "Temperature Probe",
        29: "Electrical Current Probe",
        30: "Out-of-band Remote Access",
        31: "Boot Integrity Services",
        32: "System Boot",
        33: "64-bit Memory Error",
        34: "Management Device",
        35: "Management Device Component",
        36: "Management Device Threshold Data",
        37: "Memory Channel",
        38: "IPMI Device",
        39: "Power Supply",
    }

    return c
}

func DMI() {

    database, _ := db.Init()
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS dmi (id INTEGER PRIMARY KEY, component TEXT, output TEXT)")
    statement.Exec()

    statement, _ = database.Prepare("INSERT INTO dmi (id, component, output) VALUES (?, ?, ?)")

    for code, title := range CodeTable() {
        a := ["-t", strconv.Itoa(code)]
        statement.Exec(code, title, RunCommand("dmidecode", a))
    }

    rows, _ := database.Query("SELECT id, component, output FROM dmi")
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
