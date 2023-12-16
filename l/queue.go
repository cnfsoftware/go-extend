package l

// Queue represents a generic queue data structure.
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new instance of Queue with the provided items
func NewQueue[T any](items ...T) Queue[T] {
	return Queue[T]{
		items: items,
	}
}

// Length returns the number of items in the queue.
func (q *Queue[T]) Length() int {
	return len(q.items)
}

// Push appends the given item to the end of the queue.
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Pop removes and returns the first item in the queue, if there are any. If the queue is empty, it returns nil.
// The remaining items in the queue are shifted to the left by one index.
func (q *Queue[T]) Pop() *T {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return &item
}

// Peek returns a pointer to the first item in the queue.
// If the queue is empty, it returns nil.
func (q *Queue[T]) Peek() *T {
	if len(q.items) == 0 {
		return nil
	}

	return &q.items[0]
}
