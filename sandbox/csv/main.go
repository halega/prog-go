package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `first_name,last_name,username,grade
"Rob","Pike",rob
Ken,Thompson,ken,5
"Robert","Griesemer","gri",5
`
	r := csv.NewReader(strings.NewReader(in))
	r.FieldsPerRecord = -1
	r.ReuseRecord = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(record)
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
