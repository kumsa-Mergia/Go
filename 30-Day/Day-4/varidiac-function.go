package main

import "fmt"

// Variadic function that calculates the sum of any number of integers
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum:", sum(1, 2, 3))
    fmt.Println("Sum:", sum(10, 20, 30, 40))
}
