package algs

import "testing"

// isSorted returns true if a slice is sorted in ascending order.
func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func TestIsSorted(t *testing.T) {
	inputs := []struct {
		input  []int
		sorted bool
	}{
		{nil, true},
		{[]int{}, true},
		{[]int{3}, true},
		{[]int{4, -1}, false},
		{[]int{3, 4, 7, 10, 12}, true},
		{[]int{3, 4, 7, 12, 10}, false},
	}
	for _, i := range inputs {
		if actual := isSorted(i.input); actual != i.sorted {
			t.Errorf("for %v got %t, want %t", i.input, actual, i.sorted)
		}
	}
}
