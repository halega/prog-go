package main

import (
	"fmt"
	"slices"
)

type student struct {
	FirstName string
	LastName  string
	Grades    grades
}

func (s student) Name() string {
	return fmt.Sprintf("%s %s", s.LastName, s.FirstName)
}

type grades []int

func (s grades) max() int {
	return slices.Max(s)
}

func (s grades) min() int {
	return slices.Min(s)
}

func (g grades) valid() bool {
	for _, gr := range g {
		if gr < 1 || gr > 5 {
			return false
		}
	}
	return true
}

func (g grades) valid2() bool {
	return g.min() >= 1 && g.max() <= 5
}

func (s grades) sort() {
	slices.Sort(s)
}

type grades2 map[string]int

type grade struct {
	Topic string
	Grade int
}

func (g grade) Valid() bool {
	return g.Grade >= 1 && g.Grade <= 5
}

func main() {
	s := grades{1, 2, 5, 3, 4, 0}
	fmt.Println(s)
	fmt.Println("len(s) =", len(s), "s[0] =", s[0])
	fmt.Println(s[2], s[4])
	fmt.Println("max =", s.max())
	s.sort()
	fmt.Println(s)

	s2 := grades2{"Math": 5, "Programming": 5, "Literature": 5, "English": 5, "Russian": 3, "Physics": 4}
	fmt.Println(s2)
	fmt.Println("Math's grade is", s2["Math"])
}
