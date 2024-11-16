package main

import "fmt"

// Function to swap two numbers
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 20
	fmt.Println("Before swap:", x, y)

	swap(&x, &y)
	fmt.Println("After swap:", x, y)
}
