package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	idx := make(index)

	assert.Nil(t, idx.search("foo"))
	assert.Nil(t, idx.search("donut"))

	idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	assert.Nil(t, idx.search("a"))
	assert.Equal(t, idx.search("donut"), []int{1})
	assert.Equal(t, idx.search("DoNuts"), []int{1})
	assert.Equal(t, idx.search("glass"), []int{1})

	idx.add([]document{{ID: 2, Text: "donut is a donut"}})
	assert.Nil(t, idx.search("a"))
	assert.Equal(t, idx.search("donut"), []int{1, 2})
	assert.Equal(t, idx.search("DoNuts"), []int{1, 2})
	assert.Equal(t, idx.search("glass"), []int{1})
}
