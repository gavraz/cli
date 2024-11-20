package main

import (
	"context"
	"fmt"
)

func parse(commands map[string]*Command, flags []Flag, tokens []Token) (context.Context, error) {
	flagCtx := context.Background()
	for _, flag := range flags {
		flagCtx = context.WithValue(flagCtx, flag.ID(), flag)
	}

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch token.Type {
		case assignType:
			if i+2 >= len(tokens) {
				return nil, fmt.Errorf("parsing failed: missing argument for assign op")
			}
			flagToken := tokens[i+1]
			if flagToken.Type != flagType {
				return nil, fmt.Errorf("parsing failed: unexpected operand %s: expected flag type", flagToken)
			}
			valueToken := tokens[i+2]
			if valueToken.Type != valueType {
				return nil, fmt.Errorf("parsing failed: expected value type")
			}
			newFlag, err := flagCtx.Value(flagToken.Value).(Flag).Parse(valueToken.Value)
			if err != nil {
				return nil, fmt.Errorf("parsing failed: %w", err)
			}
			flagCtx = context.WithValue(flagCtx, newFlag.ID(), newFlag)
			i += 3
		case identifierType:
			if _, ok := commands[token.Value]; !ok {
				return nil, fmt.Errorf("parsing failed: unknown identifier: %s", token.Value)
			}
			i += 1
		case flagType:
			flag, ok := flagCtx.Value(token.Value).(Flag)
			if !ok {
				return nil, fmt.Errorf("parsing failed: unknown flag: %s", token.Value)
			}
			f, ok := flag.(BoolFlag)
			if ok {
				ff, _ := f.Parse("true")
				flagCtx = context.WithValue(flagCtx, flag.ID(), ff)
				i += 1
				continue
			}
			lastToken := i == len(tokens)-1
			if lastToken {
				return nil, fmt.Errorf("parsing failed: missing argument for flag")
			}
			return nil, fmt.Errorf("parsing failed: non-bool flag: %s", token.Value)
		case valueType:
			i += 1
		default:
			panic(fmt.Errorf("parsing failed: unknown token type: %s", token))
		}
	}

	return flagCtx, nil
}
