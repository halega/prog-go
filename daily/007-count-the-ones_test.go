// Daily practice
package daily

import "testing"

func countTheOnes(array [][]int) int {
	count := 0
	for i := range array {
		for j := range array[i] {
			if array[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func TestCountTheOnes(t *testing.T) {
	actual := countTheOnes([][]int{
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1},
		{1, 0},
	})
	if actual != 7 {
		t.Error("want 7, but was", actual)
	}
}
