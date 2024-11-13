"""
Practice Implementation

Try implementing a small program that uses all three control structures together. Here’s an idea for you:

Problem:

    Use a for loop to print numbers from 1 to 10.
    Use an if-else statement inside the loop to print "even" or "odd" for each number.
    Use a switch-case to categorize the number into small (1-3), medium (4-7), and large (8-10) ranges.

Solution:
"""
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        // Check if number is even or odd
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }

        // Categorize number
        switch {
        case i >= 1 && i <= 3:
            fmt.Println(i, "is small")
        case i >= 4 && i <= 7:
            fmt.Println(i, "is medium")
        case i >= 8 && i <= 10:
            fmt.Println(i, "is large")
        }
    }
}
