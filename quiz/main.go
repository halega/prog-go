package main

import (
	"fmt"
)

const logo = ` _____  _   _  _____  _____
|  _  || \ | |/  __ \|  ___|
| | | ||  \| || /  \/| |__
| | | || . '  || |    |  __|
\ \_/ /| |\  || \__/\| |___
 \___/ \_| \_/ \____/\____/`

func main() {
	fmt.Println(logo)
	fmt.Println()

	fmt.Println("Enter your domain:")
	var domain string
	fmt.Scan(&domain)
	fmt.Printf("Your domain is %s. Correct? [y/n] ", domain)
	var answer string
	fmt.Scan(&answer)
	if answer == "y" {
		fmt.Println("Succeed")
	}
}
