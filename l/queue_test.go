package l

import (
	"go-extend/p"
	"testing"
)

func TestQueueLength(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty queue",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "nil elements",
			input:    []int{0, 0, 0},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.input...)
			got := q.Length()

			if got != tc.expected {
				t.Errorf("unexpected length - want: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	testCases := []struct {
		name        string
		input       []string
		expectedLen int
	}{
		{
			name:        "push single",
			input:       []string{"first"},
			expectedLen: 1,
		},
		{
			name:        "push multiple",
			input:       []string{"first", "second", "third"},
			expectedLen: 3,
		},
		{
			name:        "push empty",
			input:       []string{},
			expectedLen: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var q Queue[string]
			for _, item := range tc.input {
				q.Push(item)
			}
			if len := q.Length(); len != tc.expectedLen {
				t.Errorf("expected %d, got %d", tc.expectedLen, len)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name  string
		in    []int
		want  *int
		after []int
	}{
		{
			name:  "empty queue",
			in:    []int{},
			want:  nil,
			after: []int{},
		},
		{
			name:  "one item",
			in:    []int{1},
			want:  func() *int { i := 1; return &i }(),
			after: []int{},
		},
		{
			name:  "multiple items",
			in:    []int{1, 2, 3},
			want:  func() *int { i := 1; return &i }(),
			after: []int{2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue[int](tc.in...)

			got := q.Pop()
			if (got == nil && tc.want != nil) || (got != nil && *got != *tc.want) {
				t.Errorf("Pop() = %v, want %v", got, tc.want)
			}

			if len(q.items) != len(tc.after) {
				t.Errorf("Leftover Queue length - got: %d, want: %d", len(q.items), len(tc.after))
			} else {
				for i := range tc.after {
					if q.items[i] != tc.after[i] {
						t.Errorf("Leftover Queue item %d - Got: %v, Want: %v ", i, q.items[i], tc.after[i])
					}
				}
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type queueItem struct {
		queue Queue[int]
		want  *int
	}

	testCases := []struct {
		name string
		data queueItem
	}{
		{
			name: "non empty queue",
			data: queueItem{
				queue: NewQueue[int](1, 2, 3),
				want:  p.Ptr(1),
			},
		},
		{
			name: "MultiplePeek",
			data: queueItem{
				queue: NewQueue[int](1, 2, 3),
				want:  p.Ptr(1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.data.queue.Peek()
			if !p.Equal(got, tc.data.want) {
				t.Errorf("reject Peek() = %v, want %v", got, tc.data.want)
			}
		})

		t.Run(tc.name+"_AfterMultiplePeek", func(t *testing.T) {
			tc.data.queue.Peek()
			got := tc.data.queue.Peek()
			if !p.Equal(got, tc.data.want) {
				t.Errorf("reject Peek() = %v, want %v", got, tc.data.want)
			}
		})
	}
}
