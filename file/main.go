package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	count(input)
	fields(input)
	utf8symbols(input)
	words(input)
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
	words := strings.Split("one two  three\tfour\nfive", " ")
	fmt.Printf("words = %q\n", words)
}
