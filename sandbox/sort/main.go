package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Russ", "Alan", "Robert", "Niklaus"}
	slices.Sort(names)
	fmt.Println(names)
}
