package main

import (
	"context"
	"fmt"
)

func parse(commands map[string]*Command, flags []Flag, tokens []Token) (context.Context, error) {
	requiredFlags := map[string]Flag{}
	flagCtx := context.Background()
	for _, flag := range flags {
		flagCtx = context.WithValue(flagCtx, flag.ID(), flag)
		if flag.Obligatory() {
			requiredFlags[flag.ID()] = flag
		}
	}

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch token.Type {
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
			if i+1 >= len(tokens) {
				return nil, fmt.Errorf("parsing failed: missing value for flag: %s", flag.ID())
			}
			flagValueToken := tokens[i+1]
			flagWithValue, err := flag.WithValue(flagValueToken.Value)
			if err != nil {
				return nil, fmt.Errorf("parsing failed: %w", err)
			}
			flagCtx = context.WithValue(flagCtx, flagWithValue.ID(), flagWithValue)
			delete(requiredFlags, flagWithValue.ID())
			i += 2
			// TODO
			//		3. Add UT for these cases
			//		4. Support --help
		case valueType:
			return nil, fmt.Errorf("parsing failed: unexpected value type: %s", token.Value)
		default:
			panic(fmt.Errorf("parsing failed: unknown token type: %s", token))
		}
	}

	if len(requiredFlags) > 0 {
		return nil, fmt.Errorf("parsing failed: missing required flags %v", requiredFlags)
	}

	return flagCtx, nil
}
