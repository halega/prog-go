package main

import (
	"testing"

	"github.com/google/uuid"
)

type subscription struct {
	topicId uuid.UUID
	attrs   string
}

type testData struct {
	s        subscription
	httpCode int
	err      string
}

func TestScenarios(t *testing.T) {
	scenarios := map[string]testData{
		"Отправлен запрос на создание подписки несуществующего топика": func() testData {
			fakeTopicID := uuid.New()
			return testData{subscription{fakeTopicID, "R|RW"}, 404, "RESOURCE_NOT_FOUND"}
		}(),
		"Отправлен запрос на создание подписки c пустыми атрибутами": {subscription{uuid.New(), "R"}, 500, "ATTR_NOT_FOUND"},
	}
	for scenario, data := range scenarios {
		t.Log(scenario, data)
	}
	t.Fail()
}
