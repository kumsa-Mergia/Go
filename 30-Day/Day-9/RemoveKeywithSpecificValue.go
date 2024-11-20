package main

import "fmt"

func removeEntries(m map[string]int, valueToRemove int) {
    for key, value := range m {
        if value == valueToRemove {
            delete(m, key)
        }
    }
}

func main() {
    m := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 25}
    fmt.Println("Before:", m)
    removeEntries(m, 25)
    fmt.Println("After:", m)
}
