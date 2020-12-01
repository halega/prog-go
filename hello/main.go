package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "QT") {
			fmt.Println(env)
		}
	}
}
