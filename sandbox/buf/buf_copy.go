package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("buf_copy.go")
	defer file.Close()
	var buf bytes.Buffer
	buf.ReadFrom(file)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
}
