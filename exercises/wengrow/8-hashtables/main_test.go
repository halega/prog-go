package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		arrA   []int
		arrB   []int
		result []int
	}{
		{nil, nil, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{0}, []int{1}, []int{}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{0}, []int{0, 0}, []int{0}},
		{[]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6}, []int{2, 4}},
	}
	for _, tc := range testCases {
		if r := intersection(tc.arrA, tc.arrB); !cmp.Equal(r, tc.result) {
			t.Errorf(cmp.Diff(r, tc.result))
		}
	}
}
