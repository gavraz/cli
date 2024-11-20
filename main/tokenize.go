package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tokenType int

const (
	assign = iota
	identifier
	flag
	value
)

type Token struct {
	fmt.Stringer
	Type  tokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case assign:
		return t.Value
	case identifier:
		return t.Value
	case flag:
		return t.Value
	case value:
		return t.Value
	default:
		return "unknown token type"
	}
}

func tokenize(input []string) (tokens []Token) {
	for _, val := range input {
		if before, after, found := strings.Cut(val, "="); found {
			if strings.HasPrefix(before, "--") {
				tokens = append(tokens, Token{Type: flag, Value: before[2:]})
			} else {
				tokens = append(tokens, Token{Type: identifier, Value: before})
			}
			tokens = append(tokens, Token{Type: assign, Value: "="})
			tokens = append(tokens, Token{Type: value, Value: after})
			continue
		}
		if strings.HasPrefix(val, "--") {
			tokens = append(tokens, Token{Type: flag, Value: val[2:]})
			continue
		}
		if _, err := strconv.Atoi(val); err == nil {
			tokens = append(tokens, Token{Type: value, Value: val})
			continue
		}
		if _, err := strconv.ParseFloat(val, 32); err == nil {
			tokens = append(tokens, Token{Type: value, Value: val})
			continue
		}
		tokens = append(tokens, Token{Type: identifier, Value: val})
	}

	return tokens
}
