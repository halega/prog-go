package main

type reader interface {
	Read() string
}

type nop struct {
	reader
}

func main() {
	n := nop{}
	n.Read()
}
