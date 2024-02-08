package main

import (
	"errors"
	"fmt"
)

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() int {
	return i.A * 2
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func do1() {
	errors.New("Hello")
	fmt.Errorf("Hello")
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.IntPrinter(o.Double()))
	fmt.Println(o.Inner.IntPrinter(o.Double()))
}
