package main

import "fmt"

// Function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
    fahrenheit := (celsius * 9/5) + 32
    return fahrenheit
}

func main() {
    fmt.Println("30°C in Fahrenheit:", celsiusToFahrenheit(30))
    fmt.Println("100°C in Fahrenheit:", celsiusToFahrenheit(100))
}
