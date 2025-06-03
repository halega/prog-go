package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

var helloList = []string{
	"Hello, world",
	"Здравствуй, Мир!",
	"こんにちは世界",
	"Καλημέρα κόσμε",
}

func main() {
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
