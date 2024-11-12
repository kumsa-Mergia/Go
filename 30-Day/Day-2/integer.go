/*
Data Types

	Integers: int, int8, int16, int32, int64, uint (unsigned integers)
*/
package main

import "fmt"

func main() {
	var age int = 30
	var smallNumber int8 = -128
	var bigNumber int64 = 9223372036854775807

	fmt.Println("Age:", age)
	fmt.Println("Small Number (int8):", smallNumber)
	fmt.Println("Big Number (int64):", bigNumber)
}
