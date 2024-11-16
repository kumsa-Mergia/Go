package main

import "fmt"

func safeDivision(a, b int) {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}

	fmt.Println("Result:", a/b)
}

func main() {

	safeDivision(10, 2)
	safeDivision(10, 0)
	fmt.Println("Program continues...")
}
