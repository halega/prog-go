package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("checkFile: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := checkFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}
