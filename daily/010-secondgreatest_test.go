package daily

import "testing"

// secondGreatest searches for second greatest value index
// without sorting array
func secondGreatest(arr []int) int {
	g, sg := 0, -1
	for i := 1; i < len(arr); i++ {
		if arr[g] < arr[i] {
			g, sg = i, g
		} else if sg == -1 || arr[sg] < arr[i] && arr[g] != arr[i] {
			sg = i
		}
	}
	return sg
}

func TestSecondGreatest(t *testing.T) {
	cases := []struct {
		arr []int
		r   int
	}{
		{nil, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 1}, 1},
		{[]int{3, 1, 2}, 2},
		{[]int{1, 2, 3}, 1},
		{[]int{2, 3, 1}, 0},
		{[]int{1, 2, 2}, 0},
		{[]int{1, 2, 2, 2}, 0},
		{[]int{1, 1, 2}, 0},
		{[]int{2, 1, 1}, 1},
		{[]int{2, 2, 1, 1}, 1},
		{[]int{2, 1, 1, 1}, 1},
	}
	for _, c := range cases {
		if r := secondGreatest(c.arr); r != c.r {
			t.Errorf("%v: want %d but got %d", c.arr, c.r, r)
		}
	}
}
