package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Function to update the age of a person
func updateAge(p *Person, newAge int) {
	p.Age = newAge // Modify the struct field through the pointer
}

func main() {
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Before update:", person)

	updateAge(&person, 30) // Pass the address of the struct
	fmt.Println("After update:", person)
}
