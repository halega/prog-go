package main

import (
	"fmt"
	"math/bits"
)

func main() {
	const c = 1.85e19
	cc := 9_223_372_036_854_775_807
	ff := int(9.223e18)
	var x int
	var f float64
	var t int = 11
	var tt int = 1_000
	var b int = 0b11
	var o int = 0o11
	var oo int = 0_11
	var h int = 0x11
	var is bool

	fmt.Println(c, cc, ff, x, f, t, tt, b, o, oo, h, is)
	fmt.Println(inc(1))

	fmt.Println(bits.Len(10))
	divide()

	for i := range 10 {
		fmt.Println(i)
	}

	var i int
	var sum int
	for i = range 50 {
		sum += i
	}
	fmt.Println(sum)
}

func inc(n uint8) byte {
	var _1_ byte = 1
	return n + _1_
}

func divide() {
	x := 5
	y := 2
	fmt.Println(float64(x) / float64(y))
}
