// Daily practice
package daily

import "testing"

func buildWords(chars []string) []string {
	var words []string
	for _, c1 := range chars {
		for _, c2 := range chars {
			words = append(words, c1+c2)
		}
	}
	return words
}

func TestBuildWords(t *testing.T) {
	t.Log(buildWords([]string{"a", "b", "c", "d"}))
}
