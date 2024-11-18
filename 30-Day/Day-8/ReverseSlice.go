package main

import "fmt"

func reverse(slice []int) []int {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
    return slice
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Original:", nums)
    fmt.Println("Reversed:", reverse(nums))
}
