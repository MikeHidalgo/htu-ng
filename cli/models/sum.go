package models

import (
    "htu-ng/cli/db"
)

func Summary() {

    database, _ := db.DB("./htu.db")
}
