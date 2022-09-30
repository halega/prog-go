package main

// firstDuplication returns the fisrt duplicate value found in a string slice.
//
// 2. Write a function that accepts an array of strings and returns the first
// duplicate value it finds. For example, if the array is ["a", "b", "c", "d", "c", "e",
// "f"], the function should return "c", since it’s duplicated within the array.
// (You can assume that there’s one pair of duplicates within the array.)
// Make sure the function has an efficiency of O(N).
func firstDuplicate(input []string) string {
	m := map[string]bool{}
	for _, s := range input {
		if m[s] {
			return s
		}
		m[s] = true
	}
	return ""
}
