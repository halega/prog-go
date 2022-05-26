// Daily practice 26.05.2022.
package daily

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// SelectionSort sorts the input in ascending order.
func SelectionSort1(input []int) {
	for i := 0; i < len(input)-1; i++ {
		minValueIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minValueIndex] {
				minValueIndex = j
			}
		}
		if minValueIndex != i {
			input[i], input[minValueIndex] = input[minValueIndex], input[i]
		}
	}
}

func TestSelectionSort1(t *testing.T) {
	input := []int{1, 0, -1, 4, 2, 2, 10, 3, 4}
	expected := []int{-1, 0, 1, 2, 2, 3, 4, 4, 10}
	SelectionSort1(input)
	if !cmp.Equal(expected, input) {
		t.Error("want", expected, "but was", input)
	}
}
