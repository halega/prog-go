package main

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const resultS = `{
	"id": 1,
	"name": "Agent Emulator",
	"url": "https://jira.com/agent",
	"createdAt": "2020-11-30T17:36:00Z",
	"tasks": [
		{
			"taskId": 1,
			"url": "https://jira.com/agent/1"
		},
		{
			"taskId": 3,
			"url": "https://jira.com/agent/3"
		},
		{
			"taskId": 5,
			"url": "https://jira.com/agent/5"
		},
		{
			"taskId": 2,
			"url": "https://jira.com/agent/2"
		}
	]
}`

type jsonStruct struct {
	ID        int
	Name      string
	URL       string
	CreatedAt string
	Tasks     []struct {
		TaskID int
		URL    string
	}
}

func TestAll(t *testing.T) {
	funcs := []struct {
		f    func() ([]byte, error)
		name string
	}{
		{f: byStd, name: "byStd"},
		{f: byMapStructure, name: "byMapStructure"},
		{f: byMapStructurePartial, name: "byMapStructurePartial"},
		{f: bySjson, name: "bySjson"},
	}

	want := unmarshal([]byte(resultS))
	for _, f := range funcs {
		newS, err := f.f()
		if err != nil {
			t.Errorf("Error on %s(): %s", f.name, err)
		}
		got := unmarshal(newS)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("%s() mismatch (-want +got):\n%s", f.name, diff)
		}

	}

}

func unmarshal(data []byte) *jsonStruct {
	var s jsonStruct
	if err := json.Unmarshal(data, &s); err != nil {
		panic(err)
	}
	return &s
}

func BenchmarkByStd(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byStd()
	}
	_ = result
}

func BenchmarkByMapStructure(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byMapStructure()
	}
	_ = result
}

func BenchmarkByMapStructurePartial(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = byMapStructurePartial()
	}
	_ = result
}

func BenchmarkBySjson(b *testing.B) {
	var result []byte
	for i := 0; i < b.N; i++ {
		result, _ = bySjson()
	}
	_ = result
}
