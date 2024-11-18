package main

import "fmt"

type Car struct {
    Brand string
    Model string
    Year  int
}

func (c Car) DisplayDetails() {
    fmt.Printf("Car: %s %s (%d)\n", c.Brand, c.Model, c.Year)
}

func (c *Car) UpdateYear(newYear int) {
    c.Year = newYear
}

func main() {
    car := Car{Brand: "Toyota", Model: "Corolla", Year: 2015}
    car.DisplayDetails()

    car.UpdateYear(2023)
    car.DisplayDetails()
}
