package main

import "fmt"

func main() {
	var p *int // Declare a nil pointer

	if p == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}
}
