package main

import (
    "fmt"
    "time"
)

func ping(ch chan string) {
    for {
        ch <- "Ping"
        time.Sleep(time.Second)
    }
}

func pong(ch chan string) {
    for {
        fmt.Println(<-ch) // Print "Ping"
        fmt.Println("Pong")
        time.Sleep(time.Second)
    }
}

func main() {
    ch := make(chan string)
    go ping(ch)
    go pong(ch)

    time.Sleep(5 * time.Second) // Let the game run for 5 seconds
}
