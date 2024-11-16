package main

import "fmt"

// Function to append a value to a slice

func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value) // Dereference and update
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Before append:", nums)

	appendValue(&nums, 4)
	fmt.Println("After append:", nums)
}
