# generic-set-go
In Go, there is no built-in "set" type, but you can implement a set-like structure using `map[T]struct{}`, since ``struct{}`` takes zero bytes.

## Key Features

- ✅ Works with any comparable type (int, string, etc.)
- ✅ No duplicates allowed
- ✅ Minimal memory overhead
- ✅ Fast lookups (O(1)) using a map[T]struct{}
- ❌ Order is NOT preserved

## Why the Order is NOT preserved?

In Go, maps are **unordered collections**. This means that when you iterate over a map, the order of keys **is not guaranteed** to be the same as the order in which they were inserted.

## Guide

### Instalation

```sh
// go get github.com/leoapsilva/generic-set-go
go get github.com/leoapsilva/generic-set-go
``` 

### Example

Quick usage example:

```go
package main

import (
    "fmt"
    set "github.com/leoapsilva/generic-set-go"
)

func main() {
	// Integer Set
	intSet := set.New[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(10) // Duplicate, won't be added
	fmt.Println("Integer Set:", intSet.Values()) // [10 20]

	// String Set
	stringSet := set.New[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")
	stringSet.Add("apple") // Duplicate, won't be added
	fmt.Println("String Set:", stringSet.Values()) // ["apple" "banana"]
}
``` 

## License
[MIT](LICENSE)