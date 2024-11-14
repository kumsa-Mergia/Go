cover the three main types: if-else, switch-case, and for loops. These are fundamental building blocks for managing the flow of your program.
## 1. If-Else Statements

The if-else statement is used to execute a block of code based on a condition.

Syntax:
```go
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
```

You can also chain multiple else if conditions.

## 2. Switch-Case Statements

switch is an alternative to multiple if-else conditions. It evaluates an expression and matches it with cases. It’s cleaner and more readable when there are many conditions to check.

Syntax:
```go
switch expression {
case value1:
    // code to execute if expression == value1
case value2:
    // code to execute if expression == value2
default:
    // code to execute if no case matches
}
```

## 3. For Loop

The for loop in Go is the only loop construct. It can be used in three ways: as a traditional for loop (with initialization, condition, and post statement), as a while loop, or as an infinite loop.

Syntax:

// Traditional for loop
```go

for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
// While loop (no initialization or post statement)

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```
Syntax:
// Infinite loop
```go

for {
    fmt.Println("This runs forever!")
    break // To break the loop, otherwise it will run infinitely
}
```


'''
