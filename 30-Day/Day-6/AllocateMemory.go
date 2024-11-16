package main

import "fmt"

func main() {
	p := new(int)                        // Allocate memory for an integer
	fmt.Println("Value at pointer:", *p) // Default value is 0

	*p = 42
	fmt.Println("Updated value:", *p)
}
