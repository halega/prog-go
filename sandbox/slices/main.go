package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("s1:", len(s1), cap(s1))
	s2 := s1[1:]
	fmt.Println("s2:", len(s2), cap(s2))
}
