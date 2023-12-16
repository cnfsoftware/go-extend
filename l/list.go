package l

import (
	"reflect"
	"sync"
)

// List represents a generic list data structure.
type List[T any] struct {
	items []T
}

// NewList creates a new instance of the List struct with an empty items slice.
func NewList[T any](items ...T) List[T] {
	return List[T]{items: items}
}

// Add appends items to the list.
// It takes a variadic parameter `item` of type `T` and adds them to the `items` slice of the list.
// Example usage:
// l := List[int]{items: []int{1, 2, 3}}
// l.Add(4, 5, 6) -> l.items will be []int{1, 2, 3, 4, 5, 6}
// l.Add([]int{7, 8}...) -> l.items will be []int{1, 2, 3, 4, 5, 6, 7, 8}
func (l *List[T]) Add(item ...T) {
	l.items = append(l.items, item...)
}

// Get returns a pointer to the item at the specified index in the list.
// It takes an integer parameter `index` that represents the index of the item to be returned.
// Example usage:
// l := List[int]{items: []int{1, 2, 3}}
// item := l.Get(1) -> item will be a pointer to the integer value 2
// To access the value at the pointer, you can use * operator, e.g., *item will be 2.
// Note that modifying the value through the pointer will also modify the value in the list.
// It's important to keep in mind the index range to avoid index out of bounds errors.
func (l *List[T]) Get(index int) *T {
	return &l.items[index]
}

// Slice returns a slice containing all the items in the list.
// It does not modify the list itself.
// Usage example:
// list := NewList[int]()
// list.Add(1, 2, 3)
// result := list.Slice() // result will be []int{1, 2, 3}
func (l *List[T]) Slice() []T {
	if l.IsEmpty() {
		return []T{}
	}

	return l.items
}

// Insert inserts an item at the specified index in the list.
// It takes an integer parameter `index` indicating the position where the item should be inserted.
// It also takes a parameter `item` of type `T` that represents the item to be inserted.
// The method appends the item to the slice `items` at the specified index, shifting the existing elements to the right.
//
// Example usage:
// l := NewList[int]()
// l.Insert(0, 1) -> l.items will be []int{1}
// l.Insert(1, 2) -> l.items will be []int{1, 2}
// l.Insert(1, 3) -> l.items will be []int{1, 3, 2}
// l.Insert(0, 4) -> l.items will be []int{4, 1, 3, 2}
//
// Note: The index parameter should be a non-negative integer. If the index is greater than the length of the slice,
// the item will be appended to the end of the list.
func (l *List[T]) Insert(index int, item T) {
	l.items = append(l.items[:index], append([]T{item}, l.items[index:]...)...)
}

// IsEmpty returns true if the list is empty, false otherwise.
func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

// Length returns the number of items in the list.
// It takes no parameters and returns an integer representing the length of the list.
// Example usage:
// l := List[int]{items: []int{1, 2, 3, 4, 5}}
// length := l.Length() -> length will be 5
func (l *List[T]) Length() int {
	return len(l.items)
}

// ForEach applies a function `f` to each item in the list.
// It takes a function `f` of type `func(index int, item T)`, where `index` is the index of the item in the list and `item` is the item itself.
// The function `f` is called for each item in the list, with the index and item as arguments.
//
// Example usage:
// l := NewList[int]()
// l.Add(1, 2, 3)
//
//	l.ForEach(func(index int, item int) {
//	    fmt.Printf("Item at index %d: %d\n", index, item)
//	})
//
// Output:
// Item at index 0: 1
// Item at index 1: 2
// Item at index 2: 3
func (l *List[T]) ForEach(f func(index int, item T)) {
	for index, item := range l.items {
		f(index, item)
	}
}

// ParallelForEach takes a function `f` as a parameter which is executed for each item in the list concurrently.
// The function `f` receives two parameters: an index (int) and an item (T) from the list.
// The `sync.WaitGroup` `wg` is used to ensure all goroutines finish execution before returning.
// For each item in the list, a goroutine is created which calls the function `f` with the index and item as arguments.
// At the end of each goroutine, `wg.Done()` is called to indicate that the goroutine has finished execution.
// After creating all goroutines, `wg.Wait()` is called to wait for all goroutines to complete execution.
func (l *List[T]) ParallelForEach(f func(index int, item T)) {
	wg := sync.WaitGroup{}
	wg.Add(len(l.items))
	for index, item := range l.items {
		go func(index int, item T) {
			defer wg.Done()
			f(index, item)
		}(index, item)
	}
	wg.Wait()
}

// Remove removes an item from the list at the specified index.
// If the index is less than 0 or greater than or equal to the length of the list, no action is taken.
// The items after the removed item are shifted down to fill the gap.
// Example usage:
// lst := List[int]{items: []int{1, 2, 3, 4, 5}}
// lst.Remove(2) -> lst.items will be []int{1, 2, 4, 5}
// lst.Remove(0) -> lst.items will be []int{2, 4, 5}
// lst.Remove(4) -> lst.items will be []int{2, 4, 5}
// lst.Remove(5) -> lst.items will remain []int{2, 4, 5}
func (l *List[T]) Remove(index int) {
	if index < 0 || index >= len(l.items) {
		return
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
}

// IndexOf returns the index of the first occurrence of the given item in the list.
// If the item is not found, it returns -1.
func (l *List[T]) IndexOf(item T) int {
	for i, listItem := range l.items {
		if reflect.DeepEqual(listItem, item) {
			return i
		}
	}
	return -1
}

// Contains checks whether the list contains the specified item.
// It iterates over each item in the list and uses reflect.DeepEqual
// to compare the item with each list item.
// If a match is found, it returns true, otherwise it returns false.
//
// Example usage:
// l := NewList[int]()
// l.Add(1, 2, 3, 4, 5)
// result := l.Contains(3) -> result is true
// result := l.Contains(6) -> result is false
//
// Parameters:
// - item: the item to check for in the list
//
// Returns:
// - bool: true if the list contains the item, false otherwise
func (l *List[T]) Contains(item T) bool {
	for _, listItem := range l.items {
		if reflect.DeepEqual(listItem, item) {
			return true
		}
	}
	return false
}

// Clear removes all items from the list.
// It assigns an empty slice to the `items` field of the list, effectively clearing it.
// Example usage:
// l := List[int]{items: []int{1, 2, 3, 4, 5}}
// l.Clear() -> l.items will be []int{}
// l := List[string]{items: []string{"a", "b", "c"}}
// l.Clear() -> l.items will be []string{}
// l := List[Person]{items: []Person{p1, p2, p3}}
// l.Clear() -> l.items will be []Person{}
// Note: This method does not deallocate or free up any resources held by the items in the list.
func (l *List[T]) Clear() {
	l.items = []T{}
}

// Find searches for an element in the list that satisfies the given predicate.
// It takes a findFunc parameter, which is a function that takes an element of type T and returns a boolean value.
// The function iterates over the items in the list and applies the findFunc to each element.
// If findFunc returns true for an element, the function returns the index of the element and a pointer to the element.
// If no element satisfies the findFunc, the function returns -1 and a nil pointer.
// Example usage:
// list := List[int]{items: []int{1, 2, 3}}
//
//	index, value := list.Find(func(item int) bool {
//	    return item == 2
//	})
//
// In this example, the findFunc searches for the element 2 in the list.
// The function will return index = 1 and value = &2, because the element 2 is at index 1 in the list.
func (l *List[T]) Find(findFunc func(T) bool) (int, *T) {
	for index, value := range l.items {
		if findFunc(value) {
			return index, &value
		}
	}
	return -1, nil
}

// FindAll returns a new List that contains all the items from the original List that satisfy the findFunc condition.
// It takes a findFunc parameter of type func(T) bool, which is a function that takes an item of type T and returns a boolean value.
// The function loops through each item in the original List and checks if it satisfies the findFunc condition.
// If it does, the item is added to the new List using the Add method.
// Finally, it returns a pointer to the new List.
// Example usage:
// list := NewList[int]() -> creates a new List of integers
// list.Add(1, -2, 3, -4, 5) -> list.items = []int{1, -2, 3, -4, 5}
//
//	findFunc := func(i int) bool {
//	    return i > 0
//	}
//
// foundList := list.FindAll(findFunc) -> foundList.items = []int{1, 3, 5}
func (l *List[T]) FindAll(findFunc func(T) bool) *List[T] {
	list := NewList[T]()
	for _, item := range l.items {
		if findFunc(item) {
			list.Add(item)
		}
	}
	return &list
}
