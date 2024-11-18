package main

import "fmt"

func max(slice []int) int {
    if len(slice) == 0 {
        panic("slice is empty")
    }
    maxVal := slice[0]
    for _, value := range slice {
        if value > maxVal {
            maxVal = value
        }
    }
    return maxVal
}

func main() {
    nums := []int{7, 2, 10, 4}
    fmt.Println("Max value:", max(nums))
}
