package r

import (
	"testing"
)

func TestTryCatch_Raise(t *testing.T) {
	Try(func() {
		Raise("Catch")
	}).Catch(func(a any) {
		if a != "Catch" {
			t.Errorf("Expected 'Catch', got %v", a)
		}
	})
}

func TestTryCatch_Divide0(t *testing.T) {
	Try(func() {
		a := 3
		b := a - 3
		_ = a / b
	}).Catch(func(a any) {
		if a.(error).Error() != "runtime error: integer divide by zero" {
			t.Errorf("Expected 'integer divide by zero', got '%v'", a)
		}
	})
}
