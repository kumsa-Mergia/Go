package main
import "fmt"

func main() {
    // Creating and initializing a map with names and ages
    ages := map[string]int{
        "A": 30,
        "B": 25,
        "C": 35,
    }

    // Retrieving and printing B's age
    fmt.Println("B's age:", ages["B"])
}
