// Maps: Key-value pairs (similar to dictionaries in other languages).

package main

import "fmt"

func main() {
	var capitals = map[string]string{
		"Ethiopia": "Addis Ababa",
		"USA":      "Washington, D.C.",
		"Japan":    "Tokyo",
	}

	fmt.Println("Capitals:", capitals)
	fmt.Println("Capital of USA:", capitals["USA"])
}
