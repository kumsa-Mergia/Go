package main

import "fmt"

// Function that modifies a variable using a pointer
func increment(num *int) {
	*num += 1 // Dereference the pointer and increment the value
}

func main() {
	value := 10
	fmt.Println("Before increment:", value)

	increment(&value) // Pass the address of 'value'
	fmt.Println("After increment:", value)
}
