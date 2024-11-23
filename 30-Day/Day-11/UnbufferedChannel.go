package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42 // Send value into the channel
    }()

    value := <-ch // Receive value from the channel
    fmt.Println(value) // Output: 42
}
