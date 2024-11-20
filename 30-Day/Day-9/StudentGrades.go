package main

import "fmt"

func main() {
    grades := map[string]float64{
        "Alice": 85.5,
        "Bob":   90.0,
        "Charlie": 78.0,
    }

    total := 0.0
    for _, grade := range grades {
        total += grade
    }

    avg := total / float64(len(grades))
    fmt.Printf("Average grade: %.2f\n", avg)
}
