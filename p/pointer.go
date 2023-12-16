package p

import "reflect"

func Ptr[T any](value T) *T {
	return &value
}

func IsNull(value any) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}

func Equal(a, b any) bool {
	if IsNull(a) && IsNull(b) {
		// Both pointers are nil
		return true
	}

	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	// If both values are not pointers and are equal
	if aVal.Kind() != reflect.Ptr && bVal.Kind() != reflect.Ptr && a == b {
		return true
	}

	if aVal.Kind() != reflect.Ptr || bVal.Kind() != reflect.Ptr {
		// One of the values is not a pointer
		return false
	}

	// Compare the dereferenced pointers
	return aVal.Elem().Interface() == bVal.Elem().Interface()
}
