package main

import "fmt"

func main() {
	number := 7

	switch {
	case number < 10:
		fmt.Println("Less than 10")
	case number == 10:
		fmt.Println("Equal to 10")
	default:
		fmt.Println("Greater than 10")
	}
}
