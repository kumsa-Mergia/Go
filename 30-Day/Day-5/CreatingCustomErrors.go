package main

import (
	"errors"
	"fmt"
)

func customErrorExample(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func formattedErrorExample(name string) error {
	if name == "" {
		return fmt.Errorf("invalid name: %s", name)
	}
	return nil
}

func main() {
	err := customErrorExample("")
	if err != nil {
		fmt.Println("Custom Error:", err)
	}

	err = formattedErrorExample("")
	if err != nil {
		fmt.Println("Formatted Error:", err)
	}
}
