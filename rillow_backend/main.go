package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
    search("145 Belle Meade Pl")
}

func declareDB() {
    db := map[string]string{"145 Belle Meade Pl": "4 stars", "1860 Aquarius St": "3 stars"}
}

func search(input) {
	if ok := db[input]; ok {
		fmt.Println("ok")
	}
}

