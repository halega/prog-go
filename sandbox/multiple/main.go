package main

import (
	"errors"
	"fmt"
)

func multiple() (int, string, bool, error) {
	return 0, "say", true, errors.New("no-op")
}

func main() {
	fmt.Println(multiple())
}
