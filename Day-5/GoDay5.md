# Error handling in Go

Error handling in Go is an essential part of writing robust and reliable code. Go uses a simple, explicit approach to error handling with its built-in error type. Here's how it works:

## The error Type

In Go, errors are values of the error type, which is defined as:
```go
type error interface {
    Error() string
}
```

Any value of type error must implement the Error() method, which returns a string describing the error.

## Returning Errors from Functions

A common pattern in Go is for functions to return two values: the result (if any) and an error. If the error is nil, it indicates success. Otherwise, the error value provides information about what went wrong.

## Creating Custom Errors

You can use the errors.New() function from the errors package to create a new error with a custom message. Alternatively, the fmt.Errorf() function allows you to format error messages dynamically.

## Error Wrapping and Unwrapping

In Go 1.13+, you can wrap and unwrap errors using the fmt.Errorf and errors.Unwrap functions. This is useful for preserving the context of an error.

## Handling Errors with panic and recover

For unexpected or unrecoverable errors, Go provides the panic function. However, it should be used sparingly, as Go emphasizes explicit error handling.

   - panic: Used to stop the normal execution of a program.
   - recover: Used to regain control of a program after a panic.

# Error Checking Best Practices

    1 - Always Check Errors: Handle errors explicitly whenever possible. Avoid ignoring errors by assigning them to _.
    2 - Use Custom Errors for Clarity: Create custom errors for specific scenarios to make debugging easier.
    3 - Wrap Errors for Context: When passing errors up the call stack, wrap them to provide additional context.
    4 - Avoid Using panic for Regular Errors: Use panic only for critical errors where recovery is not possible.