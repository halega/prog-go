// Program daymask shows which days are encoded in daymask.
// Each bit in a daymask represents a day.
package main

import "fmt"

const daymask = 131072

func main() {
	days := fmt.Sprintf("%b", daymask)
	for i := 1; i <= 31; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	// days only contains '0' or '1', so we can iterate by bytes (not runes)
	for i := len(days) - 1; i >= 0; i-- {
		fmt.Printf("%2c ", days[i])
	}
	fmt.Println()
}
