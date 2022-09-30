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

// firstNonDuplicateCharacter solves the exercise 8.4.
//
// 4. Write a function that returns the first non-duplicated character in a string.
// For example, the string, "minimum" has two characters that only exist
// once—the "n" and the "u", so your function should return the "n", since it
// occurs first. The function should have an efficiency of O(N).
func firstNonDuplicateCharacter(s string) string {
	// count number of character's occurrence in the s string
	characters := map[rune]int{}
	for _, c := range s {
		characters[c]++
	}
	// return the first character that is occurred only once in the s string
	// Warn: do not range over characters map because it doesn't preserve the order.
	// We need the first non duplicate character. Instead range over the original string.
	for _, c := range s {
		if characters[c] == 1 {
			return string(c)
		}
	}
	return ""
}
