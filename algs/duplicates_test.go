package algs

import "testing"

func TestHasDuplicateValueClever(t *testing.T) {
	input := []int{1, 4, 2, 7, 3, 9, 1, 5}
	if !HasDuplicateValueClever(input) {
		t.Error("input has duplicate value (1) but func returns false")
	}
}

func TestHasDuplicateValueCleverFalse(t *testing.T) {
	input := []int{1, 4, 2, 7, 3, 9, 6, 5}
	if HasDuplicateValueClever(input) {
		t.Error("input doesn't have duplicate values but func returns true")
	}
}
