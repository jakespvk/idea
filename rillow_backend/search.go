package main

import (
    "fmt"
)

func declareDB() map[string]string {
    db := map[string]string{"145 Belle Meade Pl":  "4 stars", "1860 Aquarius St": "3 stars"}
    db["636 Pacific Ave, #236"] = "2 stars"
    return db
}

func search(input string, db map[string]string) {
    if output, ok := db[input]; ok {
        fmt.Println(output)
    }
}
