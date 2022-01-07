package algs

import "testing"

func TestGreatestNumber(t *testing.T) {
	if n := GreatestNumber([]int{5, 1, 6, 4, 0}); n != 6 {
		t.Errorf("GreatestNumber returned %d, want 6", n)
	}
}
