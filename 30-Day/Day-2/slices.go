//Slices: Dynamic-sized arrays.

package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)

	// Adding a new element to the slice
	fruits = append(fruits, "orange")
	fmt.Println("Updated Fruits:", fruits)
}
