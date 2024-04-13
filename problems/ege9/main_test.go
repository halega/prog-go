package main

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		nums [4]int
		ok   bool
	}{
		{[4]int{2, 31, 30, 2}, true},
	}
	for _, test := range tests {
		if ok := valid(test.nums); ok != test.ok {
			t.Errorf("Got valid(%v) = %t, but want %t", test.nums, ok, test.ok)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		line string
		nums [4]int
		err  bool
	}{
		{"2,4,7,2", [4]int{2, 4, 7, 2}, false},
		{"2, 14, 237, 0", [4]int{2, 14, 237, 0}, false},
	}
	for _, test := range tests {
		if nums, err := parse(test.line); nums != test.nums && (err != nil) != test.err {
			t.Errorf("Got parse(%s) = %v, %s, but want %v, err != nil = %t", test.line, nums, err, test.nums, test.err)
		}
	}
}
