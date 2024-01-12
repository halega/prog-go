package main

import "fmt"

type Point struct {
	X int
	Y int
	Z int
	P []int
}

type point struct{}

func (point) SayHello() string {
	return "Hello from point"
}

func main() {
	type inte int
	var x inte = 5
	fmt.Println(point{}.SayHello(), x)
}
