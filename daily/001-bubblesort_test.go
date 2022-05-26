// Daily practice 25.05.2022.
package daily

import "testing"

// TestBubbleSort1 implements bubble sort.
// (Wengrow - A Common-Sense Guide to Data Structures and Algorithms)
func TestBubbleSort1(t *testing.T) {
	input := []int{7, 9, 3, 5, 2, 4, 4, 0, -8, 12}
	t.Log("Input:", input)
	unsortedUntilIndex := len(input) - 1
	sorted := false
	steps := 0
	for !sorted {
		steps++
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			steps++
			if input[i] > input[i+1] {
				steps++
				input[i], input[i+1] = input[i+1], input[i]
				sorted = false
			}
		}
		unsortedUntilIndex--
	}
	t.Log("Sorted:", input)
	t.Log("Steps:", steps)
}
