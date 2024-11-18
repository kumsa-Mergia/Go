package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method with a value receiver
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s, and I am %d years old.\n", p.Name, p.Age)
}

// Method with a pointer receiver
func (p *Person) Birthday() {
    p.Age += 1 // Modify the field
}

func main() {
    p := Person{Name: "Alice", Age: 25}

    // Call method with value receiver
    p.Greet()

    // Call method with pointer receiver
    p.Birthday()
    p.Greet()
}
