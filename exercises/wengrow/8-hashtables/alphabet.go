package main

import "fmt"

// missingLetter solves the exercise 8.3.
//
// 3. Write a function that accepts a string that contains all the letters of the
// alphabet except one and returns the missing letter. For example, the string,
// "the quick brown box jumps over a lazy dog" contains all the letters of the alphabet
// except the letter, "f". The function should have a time complexity of O(N).
func missingLetter(s string) string {
	alphabet := make(map[rune]bool, 26)
	for l := 'a'; l <= 'z'; l++ {
		alphabet[l] = false
	}
	for _, c := range s {
		alphabet[c] = true
	}
	fmt.Println(alphabet)
	for k, v := range alphabet {
		if !v {
			return string(k)
		}
	}
	return ""
}
