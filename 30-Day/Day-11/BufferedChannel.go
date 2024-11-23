package main

import "fmt"

func main() {
    ch := make(chan string, 2) // Create a buffered channel

    ch <- "Hello"
    ch <- "World"

    fmt.Println(<-ch) // Output: Hello
    fmt.Println(<-ch) // Output: World
}
