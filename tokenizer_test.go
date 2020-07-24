package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		text   string
		tokens []string
	}{
		{
			text:   "",
			tokens: []string{},
		},
		{
			text:   "a",
			tokens: []string{"a"},
		},
		{
			text:   "small wild,cat!",
			tokens: []string{"small", "wild", "cat"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(st *testing.T) {
			assert.EqualValues(st, tc.tokens, tokenize(tc.text))
		})
	}
}
