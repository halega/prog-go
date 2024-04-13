package main

import (
	"fmt"
	"slices"
)

const (
	_ int = iota
	StateCreated
	StateRunning
	StateCompleted
)

func main() {
	lencap()
	copyClone()
	stdClone()
	appendClone()
}

func lencap() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("s1:", len(s1), cap(s1))
	s2 := s1[1:]
	fmt.Println("s2:", len(s2), cap(s2))

	fmt.Println(StateCreated, StateRunning, StateCompleted)
	var s []int
	fmt.Println(s == nil, len(s))
}

func copyClone() {
	orig := []int{0, 1, 2, 3}
	dup := make([]int, len(orig))
	copy(dup, orig)
	fmt.Printf("  copy: orig(%d, %d): %v, dup(%d, %d): %v\n",
		len(orig), cap(orig), orig, len(dup), cap(dup), dup)
}

func stdClone() {
	orig := []int{0, 1, 2, 3}
	dup := slices.Clone(orig)
	fmt.Printf(" clone: orig(%d, %d): %v, dup(%d, %d): %v\n",
		len(orig), cap(orig), orig, len(dup), cap(dup), dup)
}

func appendClone() {
	orig := []int{0, 1, 2, 3}
	dup := append([]int{}, orig...)
	fmt.Printf("append: orig(%d, %d): %v, dup(%d, %d): %v\n",
		len(orig), cap(orig), orig, len(dup), cap(dup), dup)
}
