package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	assert.Equal(t, "1984-04-27", "1984-04-27 ", "they should be equal")
}
