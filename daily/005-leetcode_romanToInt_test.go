// Daily practice
package daily

import "testing"

// Runtime: 15 ms, faster than 36.16% of Go online submissions for Roman to Integer.
// Memory Usage: 4.8 MB, less than 5.71% of Go online submissions for Roman to Integer.
func romanToInt(s string) int {
	vals := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	decimal := 0
	for i := 0; i < len(s); i++ {
		symbol := string(s[i])
		if i+1 < len(s) && vals[symbol+string(s[i+1])] > 0 {
			decimal += vals[symbol+string(s[i+1])]
			i++
		} else {
			decimal += vals[symbol]
		}
	}
	return decimal
}

func TestRomanToInt(t *testing.T) {
	input := "MCMXCIV"
	actual := romanToInt(input)
	if actual != 1994 {
		t.Error("want", 1994, "but was", actual)
	}
}
