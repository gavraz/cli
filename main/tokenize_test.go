package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_Tokenization(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "add",
			expected: []Token{
				{Type: identifier, Value: "add"},
			},
		},
		{
			input: "--flag",
			expected: []Token{
				{Type: flag, Value: "flag"},
			},
		},
		{
			input: "git add test.go",
			expected: []Token{
				{Type: identifier, Value: "git"},
				{Type: identifier, Value: "add"},
				{Type: identifier, Value: "test.go"},
			},
		},
		{
			input: "git diff --cached",
			expected: []Token{
				{Type: identifier, Value: "git"},
				{Type: identifier, Value: "diff"},
				{Type: flag, Value: "cached"},
			},
		},
		{
			input: "123",
			expected: []Token{
				{Type: value, Value: "123"},
			},
		},
		{
			input: "var=5",
			expected: []Token{
				{Type: identifier, Value: "var"},
				{Type: assign, Value: "="},
				{Type: value, Value: "5"},
			},
		},
		{
			input: "git add var --value=10",
			expected: []Token{
				{Type: identifier, Value: "git"},
				{Type: identifier, Value: "add"},
				{Type: identifier, Value: "var"},
				{Type: flag, Value: "value"},
				{Type: assign, Value: "="},
				{Type: value, Value: "10"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := tokenize(strings.Split(test.input, " "))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("For input %q, expected tokens %v, got %v", test.input, test.expected, actual)
			}
		})
	}
}
