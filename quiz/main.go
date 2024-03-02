package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const logo = ` _____  _   _  _____  _____
|  _  || \ | |/  __ \|  ___|
| | | ||  \| || /  \/| |__
| | | || . '  || |    |  __|
\ \_/ /| |\  || \__/\| |___
 \___/ \_| \_/ \____/\____/`

func main() {
	fmt.Println(logo)
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
