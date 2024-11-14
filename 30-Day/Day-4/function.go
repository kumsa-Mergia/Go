package main

import "fmt"

// Basic function that adds two numbers
func add(x int, y int) int {
    return x + y
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result)
}
