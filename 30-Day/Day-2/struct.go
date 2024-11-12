// tructs: Custom data types, similar to classes without methods.

package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create an instance of Person
	var person = Person{Name: "Alice", Age: 30}
	fmt.Println("Person:", person)
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}
