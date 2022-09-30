package main

// intersection returns intersection between a and b slices.
// Exercise 8.1 of A Common-Sense Guide to Data Structures and Algorithms
// by Wengrow
//
// 1. Write a function that returns the intersection of two arrays. The intersec-
// tion is a third array that contains all values contained within the first two
// arrays. For example, the intersection of [1, 2, 3, 4, 5] and [0, 2, 4, 6, 8] is [2, 4].
// Your function should have a complexity of O(N). (If your programming
// language has a built-in way of doing this, don’t use it. The idea is to build
// the algorithm yourself.)
func intersection(a, b []int) []int {
	aMap := map[int]bool{}
	for _, n := range a {
		aMap[n] = true
	}
	interMap := map[int]bool{}
	for _, bVal := range b {
		if aMap[bVal] {
			interMap[bVal] = true
		}
	}
	result := make([]int, 0, len(interMap))
	for interVal := range interMap {
		result = append(result, interVal)
	}
	return result
}
