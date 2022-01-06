package algs

import "fmt"

// BubbleSort sorts nums slice in-place. Time complexity: O(N^2).
// Sample steps:
// [5, 3, 7, 1, 9]
// [3, 5, 7, 1, 9]
// [1, 5, 7, 3, 9]
// [1, 5, 7, 3, 9]
// [1, 3, 7, 5, 9]
// ...
func BubbleSort(nums []int) {
	steps := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Printf("%4d: %2v\tcompare(%2d, %2d): %2d > %2d\n", steps, nums, i, j, nums[i], nums[j])
			steps++ // 1 step for comparinson
			if nums[i] > nums[j] {
				steps++ // 1 step for swap
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
