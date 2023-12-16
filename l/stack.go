package l

type Stack[T any] struct {
	items []T
}

// NewStack creates a new instance of the Stack data structure with the specified type.
// The newly created stack is initially empty.
// Example usage:
//
//	s := NewStack[int]()
//	s.Push(10)
//	s.Push(20)
//	fmt.Println(s.Length()) // Output: 2
//
// Type parameter:
//
//	T: the type of elements that the stack will hold
//
// Returns:
//
//	Stack[T]: a new empty stack
func NewStack[T any](items ...T) Stack[T] {
	return Stack[T]{
		items: items,
	}
}

// Length returns the number of items in the stack.
func (s *Stack[T]) Length() int {
	return len(s.items)
}

// Peek returns a pointer to the top element of the stack.
func (s *Stack[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}
	item := &s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
