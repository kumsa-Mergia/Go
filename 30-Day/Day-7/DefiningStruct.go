package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a struct instance
    p := Person{Name: "Alice", Age: 25}
    fmt.Println("Person:", p)

    // Access fields
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)

    // Update fields
    p.Age = 30
    fmt.Println("Updated Age:", p.Age)
}
