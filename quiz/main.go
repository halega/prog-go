package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	var s string
	n, err := fmt.Scanf("%d, %s\n", &num, &s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Scanned items: %d, Number: %d, String: %s\n", n, num, s)
}
