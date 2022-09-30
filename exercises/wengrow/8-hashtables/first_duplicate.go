package main

// firstDuplication returns the fisrt duplicate value found in a string slice.
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
