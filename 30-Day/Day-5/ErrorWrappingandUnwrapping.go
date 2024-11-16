package main

import (
	"errors"
	"fmt"
)

func wrappedErrorExample() error {

	baseErr := errors.New("base error")
	return fmt.Errorf("wrapped error: %w", baseErr) // %w wraps the error
}

func main() {

	err := wrappedErrorExample()
	fmt.Println("Error:", err)

	// Unwrapping the error
	unwrappedErr := errors.Unwrap(err)
	fmt.Println("Unwrapped Error:", unwrappedErr)
}
