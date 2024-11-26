package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("When in doubt, fish!")
    // Keep the application running
    for {
        time.Sleep(1 * time.Hour)
    }
}