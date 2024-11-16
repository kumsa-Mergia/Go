# Understanding Pointers in Go

In Go, pointers are variables that store the memory address of another variable. Pointers allow you to directly manipulate the memory location of a variable, which can be useful for performance optimization and managing shared resources.
## 1. What is a Pointer?

    - A pointer is denoted by the * symbol.
    - The & operator is used to get the memory address of a variable.
    - The * operator (dereferencing) is used to access or modify the value at the memory address.

# Key Takeaways

    1 - Pointers in Go:
        - Use & to get the address of a variable.
        - Use * to dereference a pointer.
    2 - Function Parameters:
        - Pass pointers to functions for in-place modification.
    3 - Nil Safety:
        - Always check if a pointer is nil before dereferencing.
    4 - Memory Allocation:
        - Use new for explicit memory allocation.