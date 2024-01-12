package main

import (
	"fmt"
	"unicode/utf8"
)

// A string is in effect a read-only slice of bytes.
// A string holds arbitrary bytes:
//   const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
// The source code for the string literal is UTF-8 text.

// Go source code is always UTF-8.
// A string holds arbitrary bytes.
// A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
// Those sequences represent Unicode code points, called runes.

func main() {
	строка := "Привет"
	s := "Hello"
	fmt.Println(s)
	fmt.Println(строка)

	c := 'H'
	fmt.Printf("%v\n", c)

	fmt.Println("----------------")
	s = "Здравствуй, Мир!"
	fmt.Printf("% x\n", s)
	b := []byte(s)
	fmt.Println("len =", len(s))
	fmt.Println("bytes =", len(b))
	fmt.Println("runes =", utf8.RuneCountInString(s))
	fmt.Println("len([]rune) =", len([]rune(s)))
	return
	forLoop(s)
	fmt.Println("----------------")
	rangeLoop(s)
}

func forLoop(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %x\n", i, s[i])
	}
}

func rangeLoop(s string) {
	for i, c := range s {
		fmt.Printf("%d: %v\n", i, c)
	}
}

func cmpEqual(s1, s2 string) bool {
	return s1 == s2
}

func cmpLess(s1, s2 string) bool {
	return s1 < s2
}
