package main

import (
	"fmt"
	"strings"
)

type myString string

func (s myString) removeprefix(prefix string) myString {
	return myString(strings.TrimPrefix(string(s), prefix))
}

func (s myString) removesuffix(suffix string) myString {
	return myString(strings.TrimSuffix(string(s), suffix))
}

func print(arg ...any) {
	fmt.Println(arg...)
}

func main() {
	nostarchUrl := "https://nostarch.com/blog"
	fmt.Println(strings.TrimPrefix(nostarchUrl, "https://"))
	fmt.Println(strings.TrimSuffix(nostarchUrl, "/blog"))

	var myUrl myString = "https://oszz.ru/search"
	print(myUrl.removeprefix("https://"))
	print(myUrl.removesuffix("/search"))
}
