package main

import "fmt"

type Address struct {
    City    string
    ZipCode int
}

type Employee struct {
    Name    string
    Age     int
    Address // Embedded struct
}

func main() {
    emp := Employee{
        Name: "John",
        Age:  28,
        Address: Address{
            City:    "New York",
            ZipCode: 10001,
        },
    }

    // Access fields directly
    fmt.Println("Employee Name:", emp.Name)
    fmt.Println("City:", emp.City) // Access embedded struct field directly
}
