package algs

// InsertionSort sorts a slice of ints in-place in an ascending order
// using insertion sort algorithm.
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		t := nums[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if nums[j] > t {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = t
	}
}
