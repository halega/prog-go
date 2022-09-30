package main

import "fmt"

// intersection returns intersection between a and b slices.
// Exercise 8.1 of A Common-Sense Guide to Data Structures and Algorithms
// by Wengrow
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

func main() {
	fmt.Println(intersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8}))
}
