/*
Definition: In Go, a package is a way to organize and group related code together. Each Go file belongs to a specific package, and each package provides specific functionality.

Types of Packages:

    Main Package: The main package is special; it’s used to create executable applications. When you run a Go program, it looks for the main package and the main() function as the entry point.
    Custom Packages: These are packages you define to organize reusable code. For example, you might have package math for math utilities.

Usage: Import packages at the beginning of a Go file using the import keyword. For example:
*/

package main

import (
  "fmt"
  "math"
)
