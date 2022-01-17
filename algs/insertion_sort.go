package algs

// InsertionSort sorts a slice of ints in an ascending order
// in-place using insertion sort algorithm.
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		t := nums[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if nums[j] > t {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = t
				break
			}
		}
		// if we reach first element, place temp element there
		if j == -1 {
			nums[0] = t
		}
	}
}
