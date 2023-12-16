package p

import "testing"

func TestPtr_Var(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{
			name:  "Test Ptr With Integer",
			input: int(7),
		},
		{
			name:  "Test Ptr With String",
			input: "Hello",
		},
		{
			name:  "Test Ptr With Float",
			input: 3.14,
		},
		{
			name:  "Test Ptr With Boolean",
			input: true,
		},
		{
			name:  "TestPtrWithStruct",
			input: struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.input); *got != tt.input {
				t.Errorf("Ptr() = %v, want %v", got, tt.input)
			}
		})
	}
}

const constValue = "const value"

func TestPtr_Const(t *testing.T) {
	got := Ptr(constValue)
	if *got != constValue {
		t.Errorf("Ptr() = %v, want %v", got, constValue)
	}
}

type S struct {
	value int
}

func TestIsNull(t *testing.T) {
	var ptr *int = Ptr(1)
	var nullPtr *int
	var nullInterface interface{} = nil
	var nonNullInterface interface{} = "non-null"
	var nullStruct *S = nil
	nonNullStruct := S{1}

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{"Valid Object Pointer", ptr, false},
		{"Null Pointer", nullPtr, true},
		{"Null Interface", nullInterface, true},
		{"Non-Null Interface", nonNullInterface, false},
		{"Null Struct", nullStruct, true},
		{"Non-Null Pointer", &nonNullStruct, false},
		{"Non-Null Struct", nonNullStruct, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsNull(tc.value)
			if got != tc.want {
				t.Errorf("IsNull(%v) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    any
		b    any
		want bool
	}{
		{
			name: "null equal",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "non-null not equal",
			a:    Ptr(1),
			b:    Ptr(2),
			want: false,
		},
		{
			name: "non null equal",
			a:    Ptr(3),
			b:    Ptr(3),
			want: true,
		},
		{
			name: "non null one null",
			a:    Ptr(4),
			b:    nil,
			want: false,
		},
		{
			name: "not ptr not equal",
			a:    "Hello",
			b:    "World",
			want: false,
		},
		{
			name: "no-ptr equal",
			a:    "Hello",
			b:    "Hello",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Equal(tc.a, tc.b); got != tc.want {
				t.Errorf("Equal() = %v, want %v", got, tc.want)
			}
		})
	}
}
