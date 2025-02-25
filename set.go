// Package set provides a simple implementation of a set data structure.
// A set is a collection of unique elements, meaning that it does not allow duplicate values.
// The set is implemented using a map where the keys are the elements of the set and the values are empty structs.
// The zero value of a set is an empty set ready to use.

package set

type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates and returns a new, empty set.
// The set is implemented using a map to ensure that all elements are unique.
//
// @return *Set - a pointer to the newly created set.
func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Add inserts an element into the set.
// If the element is already present in the set, it will not be added again.
//
// @param element interface{} - The element to be added to the set.
func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// Remove deletes an element from the set.
// If the element is not present in the set, the operation has no effect.
//
// @param element interface{} - The element to be removed from the set.
func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

// Contains checks if an element is present in the set.
//
// @param element interface{} - The element to check for presence in the set.
// @return bool - Returns true if the element is found in the set, otherwise returns false.
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Clear removes all elements from the set.
// After calling this method, the set will be empty.
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Size returns the number of elements in the set.
//
// @return int - An integer representing the number of unique elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Elements returns a slice containing all elements in the set.
// The order of elements in the returned slice is not guaranteed to be consistent.
//
// @return []interface{} - A slice of interface{} containing all the elements in the set.
func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}

	return keys
}
