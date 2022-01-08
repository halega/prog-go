package algs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSelectionSort(t *testing.T) {
	input := []int{4, 1, 2, 7}
	expected := []int{1, 2, 4, 7}
	if SelectionSort(input); !cmp.Equal(input, expected) {
		t.Errorf("wrong sort order:\n%v", cmp.Diff(input, expected))
	}
}

func TestSelectionSortByIsSorted(t *testing.T) {
	inputs := [][]int{
		nil,
		{},
		{0},
		{3, 3},
		{-3, 100, 10},
		{4, 8, 18},
	}
	for _, input := range inputs {
		if SelectionSort(input); !isSorted(input) {
			t.Errorf("wrong sorting: %v", input)
		}
	}
}
