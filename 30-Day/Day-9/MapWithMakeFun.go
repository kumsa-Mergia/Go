package main
import "fmt"

func main() {
    // Creating a map using the make function
    ages := make(map[string]int)

    // Initializing the map with names and ages
    ages["A"] = 30
    ages["B"] = 25
    ages["C"] = 35

    // Retrieving and printing B's age
    fmt.Println("B's age:", ages["B"])
}
