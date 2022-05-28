// Daily practice
package daily

import "testing"

// isPalindrome checks if a word s is a palindrome.
// A palindrome is a word or phrase that reads the same both forward and backward.
func isPalindrome(s string) bool {
	leftIndex := 0
	rightIndex := len(s) - 1
	for leftIndex < len(s)/2 {
		if s[leftIndex] != s[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		s            string
		isPalindrome bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"12", false},
		{"racecar", true},
		{"kayak", true},
		{"deified", true},
	}
	for _, d := range testData {
		if isPalindrome(d.s) != d.isPalindrome {
			t.Error(d.s, "isPalindrome =", d.isPalindrome, "but was not")
		}
	}
}
