package main

import (
    "fmt"
    "strings"
)

func wordCount(input string) map[string]int {
    words := strings.Fields(input)
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

func main() {
    text := "go is fun and go is powerful"
    result := wordCount(text)
    fmt.Println(result) // Output: map[and:1 fun:1 go:2 is:2 powerful:1]
}
