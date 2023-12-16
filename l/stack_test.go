package l

import (
	"go-extend/p"
	"reflect"
	"testing"
)

func TestStack_Length(t *testing.T) {
	type TestCase struct {
		name   string
		input  Stack[int]
		output int
	}

	var tests = []TestCase{
		{
			name:   "Zero length stack",
			input:  NewStack[int](),
			output: 0,
		},
		{
			name:   "Single item stack",
			input:  Stack[int]{items: []int{1}},
			output: 1,
		},
		{
			name:   "Multiple items stack",
			input:  Stack[int]{items: []int{1, 2, 3}},
			output: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.Length()
			if tc.output != got {
				t.Errorf("Expected Length = %v; got %v", tc.output, got)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack[int]
		want  *int
	}{
		{
			name:  "empty stack",
			stack: p.Ptr(NewStack[int]()),
			want:  nil,
		},
		{
			name:  "single item stack",
			stack: &Stack[int]{items: []int{1}},
			want:  p.Ptr(1),
		},
		{
			name:  "multi item stack",
			stack: &Stack[int]{items: []int{1, 2, 3}},
			want:  p.Ptr(3),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.stack.Peek(); (got == nil && tc.want != nil) || (got != nil && tc.want == nil) || (got != nil && tc.want != nil && *got != *tc.want) {
				t.Errorf("Peek() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		stack   *Stack[int]
		want    *int
		wantErr bool
	}{
		{
			name:    "pop from empty stack",
			stack:   p.Ptr(NewStack[int]()),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "pop from non-empty stack",
			stack:   &Stack[int]{items: []int{1, 2, 3}},
			want:    p.Ptr(3),
			wantErr: false,
		},
		{
			name:    "pop from single-element stack",
			stack:   &Stack[int]{items: []int{1}},
			want:    p.Ptr(1),
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tc.wantErr {
					t.Errorf("Pop() error = %v, wantErr %v", err, tc.wantErr)
				}
			}()

			got := tc.stack.Pop()

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Pop() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		pushes []int
		want   []int
	}{
		{
			name:   "push to empty stack",
			pushes: []int{1},
			want:   []int{1},
		},
		{
			name:   "push multiple items",
			pushes: []int{1, 2, 3, 4},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "push nil",
			pushes: []int{},
			want:   []int{},
		},
	}

	// Run each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Stack
			stack := Stack[int]{}

			// Push all items to the stack
			for _, push := range tc.pushes {
				stack.Push(push)
			}

			// Check the result
			for i := 0; i < len(tc.want); i++ {
				got := stack.items[i]
				want := tc.want[i]
				if got != want {
					t.Errorf("got %v; want %v", got, tc.want[i])
				}
			}

			// Verify the length of the stack
			if stack.Length() != len(tc.pushes) {
				t.Errorf("length mismatch: got %d; want %d", stack.Length(), len(tc.pushes))
			}
		})
	}
}
