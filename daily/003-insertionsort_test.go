// Daily practice 26.05.2022.
package daily

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// InsertionSort1 sorts the input in ascending order.
func InsertionSort1(input []int) {
	for i := 1; i < len(input); i++ {
		removed := input[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if input[j] > removed {
				input[j+1] = input[j]
			} else {
				break
			}
		}
		input[j+1] = removed
	}
}

func TestInsertionSort1(t *testing.T) {
	input := []int{1, 0, -1, 4, 2, 2, 10, 3, 4}
	expected := []int{-1, 0, 1, 2, 2, 3, 4, 4, 10}
	InsertionSort1(input)
	if !cmp.Equal(expected, input) {
		t.Error("want", expected, "but was", input)
	}
}
