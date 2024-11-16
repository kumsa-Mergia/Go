package main

import "fmt"

func main() {
	a := 42 // Declare an integer variable
	p := &a // Get the pointer to the variable 'a'

	fmt.Println("Value of a:", a)         // Prints the value of 'a'
	fmt.Println("Address of a:", p)       // Prints the memory address of 'a'
	fmt.Println("Value via pointer:", *p) // Dereference the pointer to get the value

	// Modifying the value via pointer
	*p = 100
	fmt.Println("Updated value of a:", a)
}
