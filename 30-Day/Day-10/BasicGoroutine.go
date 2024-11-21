package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello() // Launch goroutine
	fmt.Println("Hello from Main!")
	time.Sleep(time.Second) // Allow goroutine to complete
}
