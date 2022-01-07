package algs

// HasDuplicateValue checks whether a slice contains duplicate values.
// Time complexity: O(N^2)
func HasDuplicateValue(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != i && nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Don't be so clever! Time complexity: O(N).
func HasDuplicateValueClever(nums []int) bool {
	// nums should contains values in range 1..10 or we'll fail
	dups := make([]bool, 10)
	for _, n := range nums {
		if dups[n-1] {
			return true
		} else {
			dups[n-1] = true
		}
	}
	return false
}
