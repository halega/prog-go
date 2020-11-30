package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
)

const s = `{
	"id": 1,
	"name": "Agent Emulator",
	"url": "https://jira.com/agent",
	"createdAt": "2020-11-30T17:36:00Z",
	"tasks": [
		{
			"taskId": 1,
			"url": "/1"
		},
		{
			"taskId": 3,
			"url": "/3"
		},
		{
			"taskId": 5,
			"url": "/5"
		},
		{
			"taskId": 2,
			"url": "/2"
		}
	]
}`

type project struct {
	Tasks []task                 `mapstructure:"tasks"`
	URL   string                 `mapstructure:"url"`
	Other map[string]interface{} `mapstructure:",remain"`
}

type task struct {
	TaskID int    `mapstructure:"taskId" json:"taskId"`
	URL    string `mapstructure:"url" json:"url"`
}

func main() {
	fmt.Printf("Source JSON:\n%s\n\n", s)

	byStd()
	byMapStructure()
	byMapStructurePartial()
	bySjson()
}

// Using only package encoding/json
func byStd() {
	fmt.Println("==== byStd")

	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)

	url := data["url"].(string)
	for _, v := range data["tasks"].([]interface{}) {
		t := v.(map[string]interface{})
		t["url"] = url + t["url"].(string)
	}

	newS, _ := json.Marshal(&data)
	fmt.Printf("%s\n", pretty.Pretty(newS))
}

// Uses package github.com/mitchellh/mapstructure.
func byMapStructure() {
	fmt.Println("==== byMapStructure")

	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)
	var p project
	mapstructure.Decode(data, &p)
	for i := range p.Tasks {
		p.Tasks[i].URL = p.URL + p.Tasks[i].URL
	}

	var dataNew map[string]interface{}
	mapstructure.Decode(p, &dataNew)

	// Include Other map into dataNew
	for k, v := range dataNew[""].(map[string]interface{}) {
		dataNew[k] = v
	}
	delete(dataNew, "")

	newS, _ := json.Marshal(&dataNew)
	fmt.Printf("%s\n", pretty.Pretty(newS))
}

// Uses package github.com/mitchellh/mapstructure on partial data.
func byMapStructurePartial() {
	fmt.Println("==== byMapStructurePartial")

	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)

	var tasks []task
	mapstructure.Decode(data["tasks"], &tasks)
	url := data["url"].(string)
	for i := range tasks {
		tasks[i].URL = url + tasks[i].URL
	}

	var tasksMap map[string]interface{}
	if err := mapstructure.Decode(&struct{ Tasks []task }{tasks}, &tasksMap); err != nil {
		panic(err)
	}

	data["tasks"] = tasksMap["Tasks"]
	newS, _ := json.Marshal(&data)
	fmt.Printf("%s\n", pretty.Pretty(newS))
}

// Uses packages github.com/tidwall/gjson and github.com/tidwall/sjson.
func bySjson() {
	fmt.Println("==== bySjson")

	projURL := gjson.Get(s, "url")
	t := gjson.Get(s, "tasks")
	var tasks []task
	if err := json.Unmarshal([]byte(t.Raw), &tasks); err != nil {
		panic(err)
	}
	for i := range tasks {
		tasks[i].URL = projURL.Str + tasks[i].URL
	}
	newS, err := sjson.Set(s, "tasks", &tasks)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", pretty.Pretty([]byte(newS)))
}
