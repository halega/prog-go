package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	assert.Equal(t, "1984-04-27", "1984-04-27 ", "they should be equal")
}

func negate(n int) int { return n }

var negateTests = []struct {
	n    int
	want int
}{
	{0, 0},
	{1, -1},
	{-1, 1},
	{10, -10},
}

func TestNegate(t *testing.T) {
	for _, test := range negateTests {
		if got := negate(test.n); got != test.want {
			t.Errorf("negate(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
