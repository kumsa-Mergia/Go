# Go Basics

## 1. Basic Go Syntax

Every Go file begins with a package declaration, followed by imports and the main function for standalone programs. Go code is organized in packages (similar to modules or libraries).

```go
package main  // declares this as the main package

import "fmt"  // imports the fmt package for formatted I/O

func main() { // the main function, where execution begins
    fmt.Println("Hello, World!")
}
```

## 2. Variables

Variables in Go are declared using the var keyword, but Go also allows implicit declaration using := in function bodies.

    Explicit Declaration:
```go
var x int = 10
var y string = "Hello"

Type Inference:

```go
var z = 20      // Go infers z as an int
var greeting = "Hi there!" // inferred as a string
```
Short Declaration:

```go
name := "Alice" // name is a string, inferred
age := 30       // age is an int, inferred
```

Multiple Variable Declaration:
```go
    var a, b, c int = 1, 2, 3
    var m, n = "hello", 45
```
## 3. Constants

Constants are values that do not change, defined using the const keyword. They must be assigned at compile-time.
```go
const Pi = 3.14
const Greeting = "Hello, Go!"
const a, b = 5, "constant string"
```

## 4. Data Types

Go has several built-in data types:
Basic Types
```go
    Integers: int, int8, int16, int32, int64 and their unsigned counterparts (uint, uint8, etc.)
    Floating Points: float32, float64
    Boolean: bool
    Strings: string
```
```go
var x int = 10         // integer
var pi float64 = 3.14  // float
var isGoFun bool = true // boolean
var language string = "Go"
```
Complex Types

    Arrays: Fixed-size list of elements of the same type.
```go
var arr [3]int = [3]int{1, 2, 3}
```
Slices: Dynamic arrays.
```go
slice := []int{1, 2, 3, 4}
```
Maps: Key-value pairs, similar to dictionaries or hash tables.
```go
var cities = map[string]string{"ETH": "Addis Ababa", "USA": "Washington"}
```
Structs: Custom data types, like classes without methods.
```go
    type Person struct {
        Name string
        Age  int
    }

    var person = Person{Name: "Alice", Age: 25}
```

Summary

    Use var or := to declare variables.
    Use const for constants.
    Understand Go's basic data types for efficient memory and type usage.
