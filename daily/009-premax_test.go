package daily

import (
	"fmt"
	"testing"
)

func preMax(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, fmt.Errorf("nums contains not enough elements")
	}
	max, preMax := nums[0], nums[1]

	for _, n := range nums {
		if n > max {
			preMax, max = max, n
		}
	}

	if preMax == max {
		return 0, fmt.Errorf("no preMax element")
	}

	return preMax, nil
}

func preMaxIndex(nums []int) (int, error) {
	if len(nums) < 2 {
		return -1, fmt.Errorf("nums contains not enough elements")
	}
	max, preMax := 0, 1

	for i := range nums {
		if nums[i] > nums[max] {
			preMax, max = max, i
		}
	}

	if nums[preMax] == nums[max] {
		return -1, fmt.Errorf("no preMax element")
	}

	return preMax, nil
}

func preMaxIndex2(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	max, preMax := 0, 1

	for i := range nums {
		if nums[max] < nums[i] {
			preMax, max = max, i
		}
	}

	if nums[preMax] == nums[max] {
		return -1
	}
	return preMax
}

func TestPreMaxIndex2(t *testing.T) {
	cases := []struct {
		nums   []int
		preMax int
	}{
		{nil, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 1}, -1},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 1},
		{[]int{2, 3, 1}, 0},
		{[]int{1, 2, 3}, 1},
		{[]int{3, 2, 1}, 1},
	}

	for _, c := range cases {
		if r := preMaxIndex2(c.nums); r != c.preMax {
			t.Errorf("%v: Want preMax = %d; got preMax = %d",
				c.nums, c.preMax, r)
		}
	}
}

func TestPreMaxIndex(t *testing.T) {
	cases := []struct {
		nums    []int
		preMax  int
		isError bool
	}{
		{nil, -1, true},
		{[]int{}, -1, true},
		{[]int{1}, -1, true},
		{[]int{1, 1}, -1, true},
		{[]int{1, 2}, 0, false},
		{[]int{2, 1}, 1, false},
		{[]int{2, 3, 1}, 0, false},
		{[]int{1, 2, 3}, 1, false},
		{[]int{3, 2, 1}, 1, false},
	}

	for _, c := range cases {
		if r, err := preMaxIndex(c.nums); c.isError != (err != nil) || r != c.preMax {
			t.Errorf("%v: Want isError = %v, preMax = %d; got error = %v, preMax = %d", c.nums, c.isError, c.preMax, err, r)
		}
	}
}

func TestPreMax(t *testing.T) {
	cases := []struct {
		nums    []int
		preMax  int
		isError bool
	}{
		{nil, 0, true},
		{[]int{}, 0, true},
		{[]int{1}, 0, true},
		{[]int{1, 1}, 0, true},
		{[]int{1, 2}, 1, false},
		{[]int{2, 1}, 1, false},
		{[]int{2, 3, 1}, 2, false},
		{[]int{1, 2, 3}, 2, false},
		{[]int{3, 2, 1}, 2, false},
	}

	for _, c := range cases {
		if r, err := preMax(c.nums); c.isError != (err != nil) || r != c.preMax {
			t.Errorf("%v: Want isError = %v, preMax = %d; got error = %v, preMax = %d", c.nums, c.isError, c.preMax, err, r)
		}
	}
}
