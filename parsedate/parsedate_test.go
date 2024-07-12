package parsedate

import (
	"testing"
	"time"
)

func TestParseTreeMsValues(t *testing.T) {
	actual, err := parse("2024-07-12T09:08:51.253Z")
	if err != nil {
		t.Error("Error: ", err)
	}
	expected := time.Date(2024, 7, 12, 9, 8, 51, int(253*time.Millisecond), time.UTC)
	if actual != expected {
		t.Errorf("Want %v, but got %v", expected, actual)
	}
}

func TestParseTwoMsValues(t *testing.T) {
	actual, err := parse("2024-07-12T09:08:51.34Z")
	if err != nil {
		t.Error("Error: ", err)
	}
	expected := time.Date(2024, 7, 12, 9, 8, 51, 340*1e6, time.UTC)
	if actual != expected {
		t.Errorf("Want %v, but got %v", expected, actual)
	}
}

func TestSub(t *testing.T) {
	created, _ := parse("2024-07-12T09:08:51.253Z")
	started, _ := parse("2024-07-12T09:08:51.34Z")
	completed, _ := parse("2024-07-12T09:25:22.477Z")
	t.Error("Full duration: ", duration(created, completed))
	t.Error("Duration: ", duration(started, completed))
	t.Error("Full duration (ms): ", duration(created, completed).Milliseconds())
	t.Error("Duration (ms): ", duration(started, completed).Milliseconds())
}
