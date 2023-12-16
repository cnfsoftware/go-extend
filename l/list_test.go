package l

import (
	"go-extend/p"
	"reflect"
	"sync"
	"testing"
)

func TestList_Add_Array(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		toAdd    []int
		expected []int
	}{
		{
			name:     "Adding to an empty list",
			initial:  []int{},
			toAdd:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Adding single element",
			initial:  []int{1, 2, 3},
			toAdd:    []int{4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Adding multiple elements",
			initial:  []int{1, 2, 3},
			toAdd:    []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Adding no elements",
			initial:  []int{1, 2, 3},
			toAdd:    []int{},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := List[int]{items: tc.initial}
			l.Add(tc.toAdd...)
			for i, item := range tc.expected {
				if *l.Get(i) != item {
					t.Errorf("Error in test %s: expected %v at index %d, got %v", tc.name, item, i, l.Get(i))
				}
			}
		})
	}
}

func TestList_Add_Single(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		toAdd    int
		expected []int
	}{
		{
			name:     "Adding to an empty list",
			initial:  []int{},
			toAdd:    1,
			expected: []int{1},
		},
		{
			name:     "Adding single element",
			initial:  []int{1, 2, 3},
			toAdd:    4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := List[int]{items: tc.initial}
			l.Add(tc.toAdd)
			for i, item := range tc.expected {
				if *l.Get(i) != item {
					t.Errorf("Error in test %s: expected %v at index %d, got %v", tc.name, item, i, l.Get(i))
				}
			}
		})
	}
}

func TestList_Slice(t *testing.T) {
	tests := []struct {
		name        string
		addElements []string
		expected    []string
	}{
		{
			name:        "No elements",
			addElements: []string{},
			expected:    []string{},
		},
		{
			name:        "Single element",
			addElements: []string{"one"},
			expected:    []string{"one"},
		},
		{
			name:        "Multiple elements",
			addElements: []string{"one", "two", "three"},
			expected:    []string{"one", "two", "three"},
		},
	}

	// Loop through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewList[string]()
			// Add elements to list
			for _, v := range tc.addElements {
				l.Add(v)
			}

			got := l.Slice()

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name        string
		initItems   []any
		insertIndex int
		insertItem  any
		expected    []any
	}{
		{
			name:        "Empty list insert at index 0",
			initItems:   []any{},
			insertIndex: 0,
			insertItem:  "item",
			expected:    []any{"item"},
		},
		{
			name:        "Non-empty list insert at index 0",
			initItems:   []any{"item1", "item2"},
			insertIndex: 0,
			insertItem:  "item0",
			expected:    []any{"item0", "item1", "item2"},
		},
		{
			name:        "Non-empty list insert at last index",
			initItems:   []any{"item1", "item2"},
			insertIndex: 2,
			insertItem:  "item3",
			expected:    []any{"item1", "item2", "item3"},
		},
		{
			name:        "Non-empty list insert at middle index",
			initItems:   []any{"item1", "item3"},
			insertIndex: 1,
			insertItem:  "item2",
			expected:    []any{"item1", "item2", "item3"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := NewList[any]()
			for _, i := range tc.initItems {
				list.Add(i)
			}
			list.Insert(tc.insertIndex, tc.insertItem)
			result := list.Slice()
			if len(result) != len(tc.expected) {
				t.Errorf("Failed %s: expected length %v, got length %v", tc.name, len(tc.expected), len(result))
				return
			}
			for i, item := range result {
				if item != tc.expected[i] {
					t.Errorf("Failed %s: expected %v at index %d, got %v", tc.name, tc.expected[i], i, item)
					break
				}
			}
		})
	}
}

func TestList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		want bool
	}{
		{
			name: "Empty list",
			list: List[int]{items: []int{}},
			want: true,
		},
		{
			name: "Non-empty list",
			list: List[int]{items: []int{1, 2, 3}},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.list.IsEmpty(); got != tc.want {
				t.Errorf("List.IsEmpty() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestList_Length(t *testing.T) {
	type test struct {
		name   string
		input  List[int]
		output int
	}

	tests := []test{
		{
			name:   "Empty list",
			input:  NewList[int](),
			output: 0,
		},
		{
			name:   "Single item list",
			input:  List[int]{items: []int{1}},
			output: 1,
		},
		{
			name:   "Multi item list",
			input:  List[int]{items: []int{1, 2, 3, 4}},
			output: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if length := tc.input.Length(); length != tc.output {
				t.Errorf("expected %d, got %d", tc.output, length)
			}
		})
	}
}

func TestListForEach(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		want []int
	}{
		{
			name: "Empty list",
			list: NewList[int](),
			want: []int{},
		},
		{
			name: "Single item list",
			list: List[int]{items: []int{1}},
			want: []int{1},
		},
		{
			name: "Multiple item list",
			list: List[int]{items: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := []int{}
			tc.list.ForEach(func(_, item int) { got = append(got, item) })

			for i, v := range got {
				if tc.want[i] != v {
					t.Errorf("Test %v failed, expected %v but got %v", tc.name, tc.want, got)
				}
			}
		})
	}
}

func TestList_ParallelForEach(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single item",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple items",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "duplicates",
			input:    []int{1, 2, 2, 1, 3},
			expected: []int{1, 2, 2, 1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := NewList[int]()
			for _, item := range tc.input {
				list.Add(item)
			}

			var actual []int
			listMux := &sync.Mutex{}

			list.ParallelForEach(func(index int, item int) {
				listMux.Lock()
				actual = append(actual, item)
				listMux.Unlock()
			})

			if len(tc.expected) != len(actual) {
				t.Fatalf("Expected length: %v, got: %v", len(tc.expected), len(actual))
			}

			// Since ParallelForEach does not guarantee order, we need to compare in an unordered manner
			actualCount := make(map[int]int)
			for _, item := range actual {
				actualCount[item]++
			}

			for _, item := range tc.expected {
				if actualCount[item] == 0 {
					t.Fatalf("Missing element in actual: %v", item)
				}
				actualCount[item]--
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	var tests = []struct {
		name         string
		list         []int
		index        int
		expectedList []int
	}{
		{
			name:         "valid index",
			list:         []int{1, 2, 3, 4, 5},
			index:        2,
			expectedList: []int{1, 2, 4, 5},
		},
		{
			name:         "first index",
			list:         []int{1, 2, 3, 4, 5},
			index:        0,
			expectedList: []int{2, 3, 4, 5},
		},
		{
			name:         "last index",
			list:         []int{1, 2, 3, 4, 5},
			index:        4,
			expectedList: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lst := List[int]{items: tc.list}
			lst.Remove(tc.index)

			var got []int
			if !lst.IsEmpty() {
				got = lst.Slice()
			}

			if !reflect.DeepEqual(got, tc.expectedList) {
				t.Errorf("expected '%v', got '%v'", tc.expectedList, got)
			}
		})
	}
}

func TestList_Contains(t *testing.T) {
	tests := []struct {
		name     string
		list     List[int]
		item     int
		expected bool
	}{
		{
			name:     "empty List",
			list:     NewList[int](),
			item:     5,
			expected: false,
		},
		{
			name:     "single item is present",
			list:     List[int]{[]int{5}},
			item:     5,
			expected: true,
		},
		{
			name:     "single item is absent",
			list:     List[int]{[]int{3}},
			item:     5,
			expected: false,
		},
		{
			name:     "multiple items is present start",
			list:     List[int]{[]int{5, 6, 7}},
			item:     5,
			expected: true,
		},
		{
			name:     "multiple items is present middle",
			list:     List[int]{[]int{5, 6, 7}},
			item:     6,
			expected: true,
		},
		{
			name:     "multiple items is present end",
			list:     List[int]{[]int{5, 6, 7}},
			item:     7,
			expected: true,
		},
		{
			name:     "multiple items is absent",
			list:     List[int]{[]int{5, 6, 7}},
			item:     8,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.list.Contains(tc.item)
			if result != tc.expected {
				t.Errorf("Contains() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestList_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		item int
		want int
	}{
		{
			name: "empty list",
			list: NewList[int](),
			item: 1,
			want: -1,
		},
		{
			name: "item exists",
			list: List[int]{items: []int{1, 2, 3}},
			item: 2,
			want: 1,
		},
		{
			name: "item does not exist",
			list: List[int]{items: []int{1, 2, 3}},
			item: 4,
			want: -1,
		},
		{
			name: "item exists multiple times",
			list: List[int]{items: []int{1, 2, 1, 3, 1}},
			item: 1,
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.IndexOf(tc.item)
			if got != tc.want {
				t.Errorf("IndexOf(%v): got %v, want %v", tc.item, got, tc.want)
			}
		})
	}
}

func TestList_Clear(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
	}{
		{
			name: "empty list",
			list: List[int]{items: []int{}},
		},
		{
			name: "single item list",
			list: List[int]{items: []int{42}},
		},
		{
			name: "multiple item list",
			list: List[int]{items: []int{1, 2, 3, 4, 5}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Clear()
			if len(tc.list.items) != 0 {
				t.Fatalf("Expected empty list after Clear(), but got: %v", tc.list.items)
			}
		})
	}
}

func TestList_Find(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		findFunc func(int) bool
		wantIdx  int
		wantVal  *int
	}{
		{
			name:  "empty list",
			items: []int{},
			findFunc: func(val int) bool {
				return val == 1
			},
			wantIdx: -1,
			wantVal: nil,
		},
		{
			name:  "not found",
			items: []int{2, 3, 4},
			findFunc: func(val int) bool {
				return val == 1
			},
			wantIdx: -1,
			wantVal: nil,
		},
		{
			name:  "find first element",
			items: []int{1, 2, 3},
			findFunc: func(val int) bool {
				return val == 1
			},
			wantIdx: 0,
			wantVal: p.Ptr(1),
		},
		{
			name:  "find middle element",
			items: []int{1, 2, 3},
			findFunc: func(val int) bool {
				return val == 2
			},
			wantIdx: 1,
			wantVal: p.Ptr(2),
		},
		{
			name:  "find last element",
			items: []int{1, 2, 3},
			findFunc: func(val int) bool {
				return val == 3
			},
			wantIdx: 2,
			wantVal: p.Ptr(3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list List[int]
			list.Add(tt.items...)
			gotIdx, gotVal := list.Find(tt.findFunc)

			if tt.wantIdx != gotIdx {
				t.Errorf("Expected index %d, but got %d", tt.wantIdx, gotIdx)
			}

			if tt.wantVal != nil {
				if gotVal == nil {
					t.Error("Expected non-nil value, but got nil")
				} else if *tt.wantVal != *gotVal {
					t.Errorf("Expected value %d, but got %d", *tt.wantVal, *gotVal)
				}
			} else if gotVal != nil {
				t.Error("Expected nil value, but got non-nil")
			}
		})
	}
}

func TestList_FindAll(t *testing.T) {
	tests := []struct {
		name     string
		list     *List[int]
		findFunc func(int) bool
		expected *List[int]
	}{
		{
			name: "PositiveNumbers",
			list: &List[int]{items: []int{1, -2, 3, -4, 5}},
			findFunc: func(i int) bool {
				return i > 0
			},
			expected: &List[int]{items: []int{1, 3, 5}},
		},
		{
			name: "EmptyInput",
			list: &List[int]{items: []int{}},
			findFunc: func(i int) bool {
				return i > 0
			},
			expected: &List[int]{items: []int{}},
		},
		{
			name: "NoMatch",
			list: &List[int]{items: []int{-1, -2, -3, -4, -5}},
			findFunc: func(i int) bool {
				return i > 0
			},
			expected: &List[int]{items: []int{}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.list.FindAll(tc.findFunc); !reflect.DeepEqual(got.Slice(), tc.expected.Slice()) {
				t.Errorf("FindAll() = %v, want %v", got.items, tc.expected.items)
			}
		})
	}
}
