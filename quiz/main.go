package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	domain := read()
	fmt.Println("Your domain is", domain)
}

func read() string {
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	return text
}

func readAll() string {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
