package main

import "fmt"

// Function with named return values
func calculate(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return // implicit return of named variables
}

func main() {
    s, d := calculate(8, 3)
    fmt.Println("Sum:", s)
    fmt.Println("Difference:", d)
}
