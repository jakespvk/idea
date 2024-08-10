package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var db = declareDB()
    fmt.Println("Hello, World!")
    for scanner.Scan() {
        input := scanner.Text()
        if input == "" {
            break
        }
        search(input, db)
    }
}

