package set

import "fmt"

func ExampleAdd_Set() {
	// Integer Set Example
	intSet := New[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(10)                                 // Duplicate, won't be added
	fmt.Println("Integer Set:", intSet.Elements()) // [10 20]
	// Output:
	// Integer Set: [10 20]
}

func ExampleRemove_Set() {
	// Integer Set Example
	intSet := New[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Remove(10)
	fmt.Println("Integer Set after removal:", intSet.Elements()) // [20]
	// Output:
	// Integer Set after removal: [20]
}

func ExampleContains_Set() {
	// Integer Set Example
	intSet := New[int]()
	intSet.Add(10)
	intSet.Add(20)
	fmt.Println("Contains 10:", intSet.Contains(10)) // true
	fmt.Println("Contains 30:", intSet.Contains(30)) // false
	// Output:
	// Contains 10: true
	// Contains 30: false
}

func ExampleSize_Set() {
	// Integer Set Example
	intSet := New[int]()
	intSet.Add(10)
	intSet.Add(20)
	fmt.Println("Size of set:", intSet.Size()) // 2
	intSet.Remove(10)
	fmt.Println("Size of set after removal:", intSet.Size()) // 1
	// Output:
	// Size of set: 2
	// Size of set after removal: 1
}

func ExampleElements_Set() {
	// Integer Set Example
	intSet := New[int]()
	intSet.Add(10)
	intSet.Add(20)
	fmt.Println("Elements of set:", intSet.Elements()) // [10 20]
	// Output:
	// Elements of set: [10 20]
}
