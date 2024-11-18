package main

import "fmt"

type Rectangle struct {
    Length, Width float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Length + r.Width)
}

func main() {
    rect := Rectangle{Length: 10, Width: 5}

    fmt.Println("Area:", rect.Area())
    fmt.Println("Perimeter:", rect.Perimeter())
}
