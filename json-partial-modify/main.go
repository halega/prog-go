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

type Project struct {
	Tasks []Task                 `mapstructure:"tasks"`
	URL   string                 `mapstructure:"url"`
	Other map[string]interface{} `mapstructure:",remain"`
}

type Task struct {
	TaskID int    `mapstructure:"taskId" json:"taskId"`
	URL    string `mapstructure:"url" json:"url"`
}

func main() {
	fmt.Printf("Source JSON:\n%s\n\n", s)

	byMapStructure()
	bySjson()
}

func byMapStructure() {
	fmt.Println("==== byMapStructure")

	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)
	var p Project
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

func bySjson() {
	fmt.Println("==== bySjson")

	projURL := gjson.Get(s, "url")
	t := gjson.Get(s, "tasks")
	var tasks []Task
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
