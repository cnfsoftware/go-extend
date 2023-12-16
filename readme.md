# Go-Extend

`Go-Extend` is a Go (Golang) library for extending the functionality of the Go language. This library provides a set of tools for making everyday programming tasks easier.

## Features

This package provides several helpful features for Go development:

### `l` Package
The "l" package in Go is a generic package for data structures, containing implementations of three fundamental data structures: lists, queues, and stacks.

1. `List` (file list.go): The List structure provides methods for working with the internal Go slice, with functions such as Add, Get, Insert, IsEmpty, Length, and ForEach.
2. `Queue` (file queue.go): The Queue is a FIFO (First-In-First-Out) data structure. It implements basic methods, such as Push (append at the end), Pop (remove from the front), Peek (check the first element), and Length (get the number of elements).
3. `Stack` (file stack.go): The Stack is a LIFO (Last-In-First-Out) data structure. It provides standard operations such as Push (append at the top), Pop (remove from the top), Peek (check the topmost element), and Length (get the number of items on the stack). 

All three structures are generic, meaning they can store any data type.

### `p` Package
The "p" package in Go focuses on pointer-related operations and provides the following functions:

1. `Ptr[T any](value T)`: This function takes a value of any type and returns a pointer to that value.
2. `IsNull(value any) bool`: This function checks if a given value (of any type) is nil or if its pointer is nil. It returns true if either condition is met.
3. `Equal(a, b any) bool`: This function compares two values of any type. It returns true if both are nil, both are not pointers and equal, or if both are pointers and point to equal values.

### `r` Package
The "r" package in Go introduces a try-catch-like error handling mechanism, which is familiar to developers from other languages, but non-native to Go. It provides the following functions:

1. `Try(try func()) *TryCatch`: This function takes a function as an argument, which encapsulates the code block to be executed. The function should take no arguments and return no values. The TryCatch instance it returns can be used to handle any panics that might occur within the try block.
2. `Catch(func(any))`: This method of the TryCatch structure takes a function as a parameter which is used to handle any panics which might have occurred in the try block. The function should take one parameter of any type.
3. `Raise(err any)`: This function immediately induces a panic with the provided error. It is used to intentionally cause a panic in the code.

## Installation

To install go-extend, run the following command:

```
go get github.com/cnfsoftware/go-extend
```

## Contributing

Any contributions in the form of suggestions and issue reports are welcome!