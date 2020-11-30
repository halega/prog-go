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

	funcs := []struct {
		f    func() ([]byte, error)
		name string
	}{
		{f: byStd, name: "byStd"},
		{f: byMapStructure, name: "byMapStructure"},
		{f: byMapStructurePartial, name: "byMapStructurePartial"},
		{f: bySjson, name: "bySjson"},
	}
	for _, f := range funcs {
		fmt.Printf("==== %s\n", f.name)
		newS, err := f.f()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", pretty.Pretty(newS))
	}
}

// Uses only standard package encoding/json.
func byStd() ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}

	url := data["url"].(string)
	for _, v := range data["tasks"].([]interface{}) {
		t := v.(map[string]interface{})
		t["url"] = url + t["url"].(string)
	}

	return json.Marshal(&data)
}

// Uses package github.com/mitchellh/mapstructure.
func byMapStructure() ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	var p project
	if err := mapstructure.Decode(data, &p); err != nil {
		return nil, err
	}
	for i := range p.Tasks {
		p.Tasks[i].URL = p.URL + p.Tasks[i].URL
	}

	var dataNew map[string]interface{}
	if err := mapstructure.Decode(p, &dataNew); err != nil {
		return nil, err
	}

	// Include Other map into dataNew
	for k, v := range dataNew[""].(map[string]interface{}) {
		dataNew[k] = v
	}
	delete(dataNew, "")

	return json.Marshal(&dataNew)
}

// Uses package github.com/mitchellh/mapstructure on partial data.
func byMapStructurePartial() ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}

	var tasks []task
	if err := mapstructure.Decode(data["tasks"], &tasks); err != nil {
		return nil, err
	}
	url := data["url"].(string)
	for i := range tasks {
		tasks[i].URL = url + tasks[i].URL
	}

	var tasksMap map[string]interface{}
	if err := mapstructure.Decode(&struct{ Tasks []task }{tasks}, &tasksMap); err != nil {
		return nil, err
	}

	data["tasks"] = tasksMap["Tasks"]
	return json.Marshal(&data)
}

// Uses packages github.com/tidwall/gjson and github.com/tidwall/sjson.
func bySjson() ([]byte, error) {
	projURL := gjson.Get(s, "url")
	t := gjson.Get(s, "tasks")
	var tasks []task
	if err := json.Unmarshal([]byte(t.Raw), &tasks); err != nil {
		return nil, err
	}
	for i := range tasks {
		tasks[i].URL = projURL.Str + tasks[i].URL
	}
	newS, err := sjson.Set(s, "tasks", &tasks)
	return []byte(newS), err
}
