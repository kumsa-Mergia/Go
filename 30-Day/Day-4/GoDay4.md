In Go, functions are a core part of the language, enabling code modularity and reusability. Let’s break down the main components of functions, including parameters, return values, and variadic functions.
# 1. Basic Functions

A function in Go is defined using the func keyword, followed by the function name, parameters (if any), return type (if any), and the function body.

Syntax:
 ```go
func functionName(parameterName parameterType) returnType {
    // function body
}
```

# 2. Parameters and Return Values

Go functions can have multiple parameters and return values. Parameters are specified as name type, and multiple parameters of the same type can be declared together.

# 3. Named Return Values

In Go, you can name return values in the function signature. This allows you to assign values directly to the return variables, and simply use return without specifying values explicitly


# 4. Variadic Functions

A variadic function accepts a variable number of arguments. It uses ... before the parameter type to signify that it can take any number of arguments of that type. Inside the function, these arguments are treated as a slice.

