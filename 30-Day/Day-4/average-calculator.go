package main

import "fmt"

// Variadic function to calculate the average of a set of numbers
func average(numbers ...float64) float64 {
    total := 0.0
    for _, num := range numbers {
        total += num
    }
    return total / float64(len(numbers))
}

func main() {
    fmt.Println("Average:", average(5.5, 3.5, 8.0))
    fmt.Println("Average:", average(10.0, 20.0, 30.0, 40.0))
}
