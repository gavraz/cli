package cli

import (
	"context"
	"fmt"
)

type parser struct {
	requiredFlags map[string]Flag
	flagCtx       context.Context
}

func initParser(flags []Flag) parser {
	p := parser{
		requiredFlags: make(map[string]Flag, len(flags)),
		flagCtx:       context.Background(),
	}

	for _, flag := range flags {
		p.flagCtx = context.WithValue(p.flagCtx, flag.ID(), flag)
		if flag.Obligatory() {
			p.requiredFlags[flag.ID()] = flag
		}
	}

	return p
}

func (p parser) Parse(tokens []Token) (context.Context, []string, error) {
	var args []string
	flagCtx := p.flagCtx

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch token.Type {
		case identifierType, valueType:
			args = append(args, token.Value)
			i += 1
		case flagType:
			flag, ok := flagCtx.Value(token.Value).(Flag)
			if !ok {
				return nil, nil, fmt.Errorf("parsing failed: unknown flag: %s", token.Value)
			}
			if i+1 >= len(tokens) {
				return nil, nil, fmt.Errorf("parsing failed: missing value for flag: %s", flag.ID())
			}
			flagValueToken := tokens[i+1]
			flagWithValue, err := flag.WithValue(flagValueToken.Value)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing failed: %w", err)
			}
			flagCtx = context.WithValue(flagCtx, flagWithValue.ID(), flagWithValue)
			delete(p.requiredFlags, flagWithValue.ID())
			i += 2
		default:
			panic(fmt.Errorf("parsing failed: unknown token type: %s", token))
		}
	}

	if len(p.requiredFlags) > 0 {
		return nil, nil, fmt.Errorf("parsing failed: missing required flags %v", p.requiredFlags)
	}

	return flagCtx, args, nil
}
