package main

import "fmt"

// Function to find the maximum number in a set of integers
func max(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    maxNum := numbers[0]
    for _, num := range numbers {
        if num > maxNum {
            maxNum = num
        }
    }
    return maxNum
}

func main() {
    fmt.Println("Max:", max(3, 5, 7, 2, 8))
    fmt.Println("Max:", max(10, 20, 15, 5))
}
