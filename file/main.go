package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "one two  three\tfour\nfive   \"quotes\""
	count(input)
	fields(input)
	utf8symbols(input)
	words(input)
	scanWords(input)
}

func count(s string) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}

func fields(s string) {
	fields := strings.Fields(s)
	fmt.Printf("strings.Fields(str) = %q\n", fields)
	fmt.Printf("len(strings.Fields(str)) = %d\n", len(fields))
}

func utf8symbols(s string) {
	chars := strings.Split(s, "")
	fmt.Printf("strings.Split(s, \"\") = %q\n", chars)
}

func words(s string) {
	words := strings.Split(s, " ")
	fmt.Printf("words = %q\n", words)
}

func scanWords(s string) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	fmt.Print("scanner = [ ")
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	fmt.Printf("]\n")
}
