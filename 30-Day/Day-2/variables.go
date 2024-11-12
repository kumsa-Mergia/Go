/*
Variables

Variables are declared with var or shorthand with :=
*/
package main

import "fmt"

func main() {
	// Explicit type declaration
	var age int = 24
	var name string = "Kumsa"

	// Type inference (Go infers the type from the assigned value)
	var greeting = "Hello, Go!"
	var temperature = 36.5 // inferred as float64

	// Short declaration (only inside functions)
	height := 169.5 // inferred as float64

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Height:", height)
}
