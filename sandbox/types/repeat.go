package main

import (
	"errors"
	"fmt"
)

type Percentage int

func (i Percentage) Module() Percentage {
	if i < 0 {
		return -i
	}
	return i
}

func Markup(price int, markup Percentage) float64 {
	return float64(price) * float64(markup) / 100.0
}

type GPoint struct {
	X int
	Y int
	Z int
}

func (p GPoint) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *GPoint) Set(x int, y int) {
	p.X = x
	p.Y = y
}

func do() {
	var x, y, z = GetCoords()
	fmt.Println(x, y, z)
}

func GetCoords() (int, error) {
	return 0, errors.New("File not found")
}

func PrintString(x *int) bool {
	*x = 5
	return true
}
