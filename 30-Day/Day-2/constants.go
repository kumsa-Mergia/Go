/*
Constants

Constants are declared using const and their values cannot be changed.
*/
package main

import "fmt"

func main() {
	// Defining constants
	const Pi = 3.14159
	const Name = "Go Language"
	const Year, Month = 2024, "November"

	fmt.Println("Pi:", Pi)
	fmt.Println("Name:", Name)
	fmt.Println("Year:", Year)
	fmt.Println("Month:", Month)
}
