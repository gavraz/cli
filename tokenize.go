package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type tokenType int

const (
	identifierType = iota
	flagType
	valueType
)

func (t tokenType) String() string {
	switch t {
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
	nextIsValue := false
	for _, val := range input {
		if nextIsValue {
			tokens = append(tokens, Token{Type: valueType, Value: val})
			nextIsValue = false
			continue
		}
		if before, after, found := strings.Cut(val, "="); found && strings.HasPrefix(before, "--") {
			tokens = append(tokens,
				Token{Type: flagType, Value: before[2:]},
				Token{Type: valueType, Value: after},
			)
			continue
		}
		if strings.HasPrefix(val, "--") {
			tokens = append(tokens, Token{Type: flagType, Value: val[2:]})
			nextIsValue = true
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
