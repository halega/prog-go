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

func d(x, y int) {

}

type Pet struct {
	Name string
	Age  int
}

func (p *Pet) GetAge() int {
	return p.Age
}

func (p *Pet) SetAge(age int) {
	p.Age = age
}

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

type Ager interface {
}

func initF(f *func()) {
	*f = func() {
		fmt.Println("inifF")
	}
}

func doF() {
	var f func()
	initF(&f)
	f()
}

func mutateP(x **int) {
	y := new(int)
	*y = 10
	*x = y
}

func doPointer() {
	x := new(int)
	*x = 5
	mutateP(&x)
	fmt.Println(*x)
}

func main() {
	do()
}
