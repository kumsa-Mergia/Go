// Arrays: Fixed-size collections of elements of the same type.

package main

import "fmt"

func main() {
	var numbers [3]int = [3]int{1, 2, 3}
	colors := [3]string{"red", "green", "blue"}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Colors:", colors)
}
