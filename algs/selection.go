package algs

// SelectionSort sorts a slice in-place in ascending order
// using selection sort with time complexity of O(N^2).
// Data movement is minimal: O(N).
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minValIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minValIndex] {
				minValIndex = j
			}
		}
		if i != minValIndex {
			nums[i], nums[minValIndex] = nums[minValIndex], nums[i]
		}
	}
}
