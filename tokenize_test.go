package cli

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
				{Type: identifierType, Value: "add"},
			},
		},
		{
			input: "--flag",
			expected: []Token{
				{Type: flagType, Value: "flag"},
			},
		},
		{
			input: "git add test.go",
			expected: []Token{
				{Type: identifierType, Value: "git"},
				{Type: identifierType, Value: "add"},
				{Type: identifierType, Value: "test.go"},
			},
		},
		{
			input: "git diff --cached",
			expected: []Token{
				{Type: identifierType, Value: "git"},
				{Type: identifierType, Value: "diff"},
				{Type: flagType, Value: "cached"},
			},
		},
		{
			input: "123",
			expected: []Token{
				{Type: valueType, Value: "123"},
			},
		},
		{
			input: "--flag=5",
			expected: []Token{
				{Type: flagType, Value: "flag"},
				{Type: valueType, Value: "5"},
			},
		},
		{
			input: "--flag 11",
			expected: []Token{
				{Type: flagType, Value: "flag"},
				{Type: valueType, Value: "11"},
			},
		},
		{
			input: "git add var --value=10",
			expected: []Token{
				{Type: identifierType, Value: "git"},
				{Type: identifierType, Value: "add"},
				{Type: identifierType, Value: "var"},
				{Type: flagType, Value: "value"},
				{Type: valueType, Value: "10"},
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
