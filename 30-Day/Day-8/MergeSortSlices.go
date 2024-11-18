package main

import (
    "fmt"
    "sort"
)

func mergeAndSort(slice1, slice2 []int) []int {
    merged := append(slice1, slice2...)
    sort.Ints(merged)
    return merged
}

func main() {
    s1 := []int{3, 1, 4}
    s2 := []int{2, 5, 6}
    result := mergeAndSort(s1, s2)
    fmt.Println("Merged and Sorted:", result)
}
