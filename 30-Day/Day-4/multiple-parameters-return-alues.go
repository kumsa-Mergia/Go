package main

import "fmt"

// Function that returns the sum and difference of two numbers
func sumAndDiff(a, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}

func main() {
    sum, diff := sumAndDiff(10, 5)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
