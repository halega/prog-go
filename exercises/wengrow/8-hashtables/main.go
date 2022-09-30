// A Common-Sense Guide to Data Structures and Algorithms by Wengrow
// Chapter 8. Solutions to the exercises.
package main

import "fmt"

func main() {
	fmt.Println("intersection() =", intersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8}))

	input := []string{"a", "b", "c", "d", "c", "e", "f"}
	fmt.Printf("firstDuplicate(%v) = '%s'\n", input, firstDuplicate(input))

	fmt.Println("missingLetter():", missingLetter("the quick brown box jumps over a lazy dog"))
	fmt.Println("firstNonDuplicateLetter():", firstNonDuplicateCharacter("minimum"))
}
