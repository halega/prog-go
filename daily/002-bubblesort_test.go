// Daily practice 26.05.2022.
package daily

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// BubbleSort2 sorts input slice in ascending order.
func BubbleSort2(input []int) {
	alreadySortedAfter := len(input) - 1
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < alreadySortedAfter; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				sorted = false
			}
		}
		alreadySortedAfter--
	}
}

func TestBubbleSort2(t *testing.T) {
	testData := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{0, -1, 1}, []int{-1, 0, 1}},
		{[]int{3, 1, -1, 4, 10, 2, 2, 4}, []int{-1, 1, 2, 2, 3, 4, 4, 10}},
	}
	for _, data := range testData {
		BubbleSort2(data.input)
		if diff := cmp.Diff(data.input, data.expected); diff != "" {
			t.Error(diff)
			t.Error(data.input, data.expected)
		}
	}
}
