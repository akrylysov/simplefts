package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowercaseFilter(t *testing.T) {
	var (
		in  = []string{"Cat", "DOG", "fish"}
		out = []string{"cat", "dog", "fish"}
	)
	assert.Equal(t, out, lowercaseFilter(in))
}

func TestStopwordFilter(t *testing.T) {
	var (
		in  = []string{"i", "am", "the", "cat"}
		out = []string{"am", "cat"}
	)
	assert.Equal(t, out, stopwordFilter(in))
}

func TestStemmerFilter(t *testing.T) {
	var (
		in  = []string{"cat", "cats", "fish", "fishing", "fished", "airline"}
		out = []string{"cat", "cat", "fish", "fish", "fish", "airlin"}
	)
	assert.Equal(t, out, stemmerFilter(in))
}
