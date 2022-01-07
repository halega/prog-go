package algs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBubbleSort4(t *testing.T) {
	nums := []int{5, 3, 7, 1}
	sorted := []int{1, 3, 5, 7}
	if t.Log(BubbleSort(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}

func TestBubbleSort8(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 4, 11, 17}
	sorted := []int{1, 3, 4, 5, 7, 9, 11, 17}
	if t.Log(BubbleSort(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}

func TestBubbleSort16(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 4, 11, 17, 2, 12, 13, 8, 0, 14, 15, -1}
	sorted := []int{-1, 0, 1, 2, 3, 4, 5, 7, 8, 9, 11, 12, 13, 14, 15, 17}
	if t.Log(BubbleSort(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}

// Tests for bubble sort from the book Wengrow.

func TestWengrowBubbleSort4(t *testing.T) {
	nums := []int{5, 3, 7, 1}
	sorted := []int{1, 3, 5, 7}
	if t.Log(WengrowBubbleSort(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}

func TestWengrowBubbleSort16(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 4, 11, 17, 2, 12, 13, 8, 0, 14, 15, -1}
	sorted := []int{-1, 0, 1, 2, 3, 4, 5, 7, 8, 9, 11, 12, 13, 14, 15, 17}
	if t.Log(WengrowBubbleSort(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}

func TestWengrowBubbleSort2(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 4, 11, 17, 2, 12, 13, 8, 0, 14, 15, -1}
	sorted := []int{-1, 0, 1, 2, 3, 4, 5, 7, 8, 9, 11, 12, 13, 14, 15, 17}
	if t.Log(WengrowBubbleSort2(nums)); !cmp.Equal(nums, sorted) {
		t.Errorf("wrong sort results; diff:\n%s", cmp.Diff(nums, sorted))
	}
}
func TestWengrowBubbleSort2WorstCase(t *testing.T) {
	input := []int{5, 4, 3, 2, 1} // sorted in descendant order
	expected := []int{1, 2, 3, 4, 5}
	if t.Log(WengrowBubbleSort2(input)); !cmp.Equal(input, expected) {
		t.Error("wrong sort result:")
		t.Error(cmp.Diff(input, expected))
	}
}
