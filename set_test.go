package set

import (
	"testing"
)

type ComplexStruct struct {
	ID   int
	Name string
}

func TestAdd(t *testing.T) {
	// Test with string set
	s := New[string]()
	s.Add("item1")
	if !s.Contains("item1") {
		t.Errorf("Expected set to contain 'item1'")
	}

	// Test with integer set
	sInt := New[int]()
	sInt.Add(1)
	if !sInt.Contains(1) {
		t.Errorf("Expected set to contain '1'")
	}

	// Test with complex struct set
	sStruct := New[ComplexStruct]()
	item := ComplexStruct{ID: 1, Name: "item1"}
	sStruct.Add(item)
	if !sStruct.Contains(item) {
		t.Errorf("Expected set to contain '%v'", item)
	}
}

func TestRemove(t *testing.T) {
	// Test with string set
	s := New[string]()
	s.Add("item1")
	s.Remove("item1")
	if s.Contains("item1") {
		t.Errorf("Expected set to not contain 'item1'")
	}

	// Test with integer set
	sInt := New[int]()
	sInt.Add(1)
	sInt.Remove(1)
	if sInt.Contains(1) {
		t.Errorf("Expected set to not contain '1'")
	}

	// Test with complex struct set
	sStruct := New[ComplexStruct]()
	item := ComplexStruct{ID: 1, Name: "item1"}
	sStruct.Add(item)
	sStruct.Remove(item)
	if sStruct.Contains(item) {
		t.Errorf("Expected set to not contain '%v'", item)
	}
}

func TestContains(t *testing.T) {
	// Test with string set
	s := New[string]()
	s.Add("item1")
	if !s.Contains("item1") {
		t.Errorf("Expected set to contain 'item1'")
	}
	if s.Contains("item2") {
		t.Errorf("Expected set to not contain 'item2'")
	}

	// Test with integer set
	sInt := New[int]()
	sInt.Add(1)
	if !sInt.Contains(1) {
		t.Errorf("Expected set to contain '1'")
	}
	if sInt.Contains(2) {
		t.Errorf("Expected set to not contain '2'")
	}

	// Test with complex struct set
	sStruct := New[ComplexStruct]()
	item1 := ComplexStruct{ID: 1, Name: "item1"}
	item2 := ComplexStruct{ID: 2, Name: "item2"}
	sStruct.Add(item1)
	if !sStruct.Contains(item1) {
		t.Errorf("Expected set to contain '%v'", item1)
	}
	if sStruct.Contains(item2) {
		t.Errorf("Expected set to not contain '%v'", item2)
	}
}

func TestSize(t *testing.T) {
	// Test with string set
	s := New[string]()
	if s.Size() != 0 {
		t.Errorf("Expected set size to be 0")
	}
	s.Add("item1")
	if s.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}
	s.Add("item2")
	if s.Size() != 2 {
		t.Errorf("Expected set size to be 2")
	}
	s.Remove("item1")
	if s.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}

	// Test with integer set
	sInt := New[int]()
	if sInt.Size() != 0 {
		t.Errorf("Expected set size to be 0")
	}
	sInt.Add(1)
	if sInt.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}
	sInt.Add(2)
	if sInt.Size() != 2 {
		t.Errorf("Expected set size to be 2")
	}
	sInt.Remove(1)
	if sInt.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}

	// Test with complex struct set
	sStruct := New[ComplexStruct]()
	if sStruct.Size() != 0 {
		t.Errorf("Expected set size to be 0")
	}
	item1 := ComplexStruct{ID: 1, Name: "item1"}
	item2 := ComplexStruct{ID: 2, Name: "item2"}
	sStruct.Add(item1)
	if sStruct.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}
	sStruct.Add(item2)
	if sStruct.Size() != 2 {
		t.Errorf("Expected set size to be 2")
	}
	sStruct.Remove(item1)
	if sStruct.Size() != 1 {
		t.Errorf("Expected set size to be 1")
	}
}

func TestElements(t *testing.T) {
	// Test with string set
	s := New[string]()
	s.Add("item1")
	s.Add("item2")
	elements := s.Elements()
	if len(elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elements))
	}
	if !s.Contains("item1") || !s.Contains("item2") {
		t.Errorf("Expected elements to contain 'item1' and 'item2'")
	}

	// Test with integer set
	sInt := New[int]()
	sInt.Add(1)
	sInt.Add(2)
	elementsInt := sInt.Elements()
	if len(elementsInt) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elementsInt))
	}
	if !sInt.Contains(1) || !sInt.Contains(2) {
		t.Errorf("Expected elements to contain 1 and 2")
	}

	// Test with complex struct set
	sStruct := New[ComplexStruct]()
	item1 := ComplexStruct{ID: 1, Name: "item1"}
	item2 := ComplexStruct{ID: 2, Name: "item2"}
	sStruct.Add(item1)
	sStruct.Add(item2)
	elementsStruct := sStruct.Elements()
	if len(elementsStruct) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elementsStruct))
	}
	if !sStruct.Contains(item1) || !sStruct.Contains(item2) {
		t.Errorf("Expected elements to contain '%v' and '%v'", item1, item2)
	}
}
