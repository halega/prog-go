package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("hello/hello")
	for _, e := range cmd.Env {
		fmt.Println(e)
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("error: cmd.Run: %s\n", err)
		os.Exit(1)
	}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: cmd.Output: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Output: %s\n", out)
}
