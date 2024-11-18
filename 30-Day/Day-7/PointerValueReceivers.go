package main

import "fmt"

type Counter struct {
    Count int
}

// Method with a value receiver
func (c Counter) Increment() {
    c.Count++ // This won't affect the original struct
}

// Method with a pointer receiver
func (c *Counter) IncrementPointer() {
    c.Count++ // This affects the original struct
}

func main() {
    c := Counter{Count: 10}

    c.Increment()
    fmt.Println("After Increment (Value Receiver):", c.Count) // Output: 10

    c.IncrementPointer()
    fmt.Println("After IncrementPointer (Pointer Receiver):", c.Count) // Output: 11
}
