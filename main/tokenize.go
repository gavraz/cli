package main

import (
	"fmt"
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

func (t tokenType) String() string {
	switch t {
	case assignType:
		return "assign"
	case identifierType:
		return "identifier"
	case flagType:
		return "flag"
	case valueType:
		return "value"
	default:
		panic("unknown token type")
	}
}

type Token struct {
	Type  tokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s:%s", t.Type, t.Value)
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
