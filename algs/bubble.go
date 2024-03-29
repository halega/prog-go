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
func BubbleSort(nums []int) string {
	trace := ""
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			steps++ // 1 step for comparinson
			trace += fmt.Sprintf("%4d: %2v\tcompare(%2d, %2d): %2d > %2d\n", steps, nums, i, j, nums[i], nums[j])
			if nums[i] > nums[j] {
				steps++ // 1 step for swap
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return trace
}

// Wengrow, A Common-Sense Guide to Data Structures and Algorithms
// Bubble Sort, Chapter 4. Speeding Up Your Code with Big O.
// Returns algorithm trace.
func WengrowBubbleSort(list []int) string {
	trace := "" // log of each algorithm step
	steps := 0
	swapped := true
	for swapped {
		swapped = false
		for i, j := 0, 1; i < len(list)-1; i, j = i+1, j+1 {
			steps++
			trace += fmt.Sprintf("%4d: %2v\tcompare(%2d, %2d): %2d > %2d %t\n", steps, list, i, j, list[i], list[j], swapped)
			if list[i] > list[j] {
				steps++
				list[i], list[j] = list[j], list[i]
				swapped = true
			}
		}
	}
	return trace
}

// WengrowBubbleSort2 sorts slice in ascendent order.
// It uses bubble sort. Time complexity: O(N^2).
func WengrowBubbleSort2(list []int) string {
	trace := "" // log of each algorithm step
	steps := 0
	unsortedUntilIndex := len(list) - 1
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			steps++
			trace += fmt.Sprintf("%4d: %2v\tcompare(%2d, %2d): %2d > %2d %t\n", steps, list, i, i+1, list[i], list[i+1], sorted)
			if list[i] > list[i+1] {
				sorted = false
				steps++
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		unsortedUntilIndex--
	}
	return trace
}
