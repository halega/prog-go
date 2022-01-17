package algs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertionSort(t *testing.T) {
	input := []int{4, 2, 7, 1, 3}
	expected := []int{1, 2, 3, 4, 7}
	if InsertionSort(input); !cmp.Equal(expected, input) {
		t.Errorf("expected %v, but was %v", expected, input)
	}
	t.Log(input)
}

func TestInsertionSortNilSlice(t *testing.T) {
	var input []int = nil
	if InsertionSort(input); !isSorted(input) {
		t.Errorf("wrong order: %v", input)
	}
	t.Log(input)
}

func TestInsertionSortSingleValue(t *testing.T) {
	input := []int{4}
	if InsertionSort(input); !isSorted(input) {
		t.Errorf("wrong order: %v", input)
	}
	t.Log(input)
}

func TestInsertionSortEqualValues(t *testing.T) {
	input := []int{4, 2, 4, 4, -1, 4}
	expected := []int{-1, 2, 4, 4, 4, 4}
	if InsertionSort(input); !cmp.Equal(expected, input) {
		t.Errorf("expected %v, but was %v", expected, input)
	}
	t.Log(input)
}
