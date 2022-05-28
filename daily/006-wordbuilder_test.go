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

func buildWordsUniq(chars []string) []string {
	var words []string
	for _, c1 := range chars {
		for _, c2 := range chars {
			if c1 != c2 {
				words = append(words, c1+c2)
			}
		}
	}
	return words
}

func buildWordsUniqIndex(chars []string) []string {
	var words []string
	for i := range chars {
		for j := range chars {
			if i != j {
				words = append(words, chars[i]+chars[j])
			}
		}
	}
	return words
}

func TestBuildWords(t *testing.T) {
	t.Log(buildWords([]string{"a", "b", "c", "d"}))
	t.Log(buildWordsUniq([]string{"a", "b", "c", "d"}))
	t.Log(buildWordsUniqIndex([]string{"a", "b", "c", "d"}))
}
