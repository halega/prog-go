package main

import "fmt"

type valuer interface {
	Value(key string) string
}

type v struct {
	key   string
	value string
}

func (x *v) Value(key string) string {
	if x.key == key {
		return x.value
	}
	return ""
}

type kv struct {
	valuer
	key   string
	value string
}

func withValue(parent valuer, key, value string) valuer {
	return &kv{parent, key, value}
}

func (x *kv) Value(key string) string {
	if x.key == key {
		return x.value
	}
	return x.valuer.Value(key)
}

func main() {
	var v1 valuer = &v{"1", "one"}
	fmt.Println(v1)
	v2 := withValue(v1, "2", "two")
	fmt.Println(v2.Value("2"))
	v3 := valuer.Value(v1, "1")
	fmt.Println(v3)
}
