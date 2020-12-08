package main

import (
	"encoding/json"
	"fmt"
)

type Obj struct {
	Ids struct {
		ID    int `json:"id"`
		NewID int `json:"newId"`
	}
	File json.RawMessage `json:"file"`
}

func main() {
	data := `{
	"Ids": {
		"id": 1,
		"newId": 100
	},
	"file" : {
		"fileName": "example.json",
		"size": 3452
	}
}`
	fmt.Printf("Old json: %s\n", data)
	j := new(Obj)
	if err := json.Unmarshal([]byte(data), j); err != nil {
		fmt.Printf("Error unmarshal JSON: %q\n", err)
	}
	j.Ids.ID, j.Ids.NewID = j.Ids.NewID, j.Ids.ID
	newData, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		fmt.Printf("Error during decode JSON: %q", err)
	}
	fmt.Printf("New JSON: %s", newData)
}
