package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Marshal (encode) to JSON
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData)) // {"name":"Alice","age":30}

	// Unmarshal (decode) from JSON
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p2) // {Name:Alice Age:30}
}
