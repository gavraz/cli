package main

import (
	"strconv"
	"strings"
)

type tokenType int

const (
	assignType = iota
	identifierType
	flagType
	valueType
)

type Token struct {
	Type  tokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case assignType:
		return t.Value
	case identifierType:
		return t.Value
	case flagType:
		return t.Value
	case valueType:
		return t.Value
	default:
		return "unknown token type"
	}
}

func tokenize(input []string) (tokens []Token) {
	for _, val := range input {
		if before, after, found := strings.Cut(val, "="); found {
			tokens = append(tokens, Token{Type: assignType, Value: "="})
			if strings.HasPrefix(before, "--") {
				tokens = append(tokens, Token{Type: flagType, Value: before[2:]})
			} else {
				tokens = append(tokens, Token{Type: identifierType, Value: before})
			}
			tokens = append(tokens, Token{Type: valueType, Value: after})
			continue
		}
		if strings.HasPrefix(val, "--") {
			tokens = append(tokens, Token{Type: flagType, Value: val[2:]})
			continue
		}
		if _, err := strconv.Atoi(val); err == nil {
			tokens = append(tokens, Token{Type: valueType, Value: val})
			continue
		}
		if _, err := strconv.ParseFloat(val, 32); err == nil {
			tokens = append(tokens, Token{Type: valueType, Value: val})
			continue
		}
		tokens = append(tokens, Token{Type: identifierType, Value: val})
	}

	return tokens
}
