package main

import (
	"fmt"
)

type App struct {
	name        string
	description string
	version     string

	commands map[string]*Command
}

func NewApp(name, description, version string) *App {
	return &App{
		name:        name,
		description: description,
		version:     version,
		commands:    make(map[string]*Command),
	}
}

func (a *App) AddCommand(command *Command) {
	a.commands[command.Name] = command
}

func (a *App) parse(tokens []Token) (*Command, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("parsing failed: missing identifier")
	}
	if tokens[0].Type != identifier {
		return nil, fmt.Errorf("parsing failed: first argument must be a identifier")
	}
	cmdName := tokens[0].Value
	cmd, ok := a.commands[cmdName]
	if !ok {
		return nil, fmt.Errorf("parsing failed: unknown identifier: %s", cmdName)
	}

	flags := make(map[string]Flag, len(cmd.Flags))
	for _, flag := range cmd.Flags {
		flags[flag.Name] = flag
	}

	const (
		noneState = iota
		commandState
		flagState
	)
	//var state = noneState
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case assign:
			return nil, fmt.Errorf("parsing failed: unexpected assign")
		case identifier:
			if _, ok := a.commands[token.Value]; !ok {
				return nil, fmt.Errorf("parsing failed: unknown identifier: %s", token.Value)
			}
		case flag:
			flag, ok := flags[token.Value]
			if !ok {
				return nil, fmt.Errorf("parsing failed: unknown flag: %s", token.Value)
			}
			lastToken := i == len(tokens)-1
			if lastToken && flag.Type != BoolFlag {
				return nil, fmt.Errorf("parsing failed: missing argument for flag")
			}
			if lastToken {
				// just a boolean flag that is true
				continue
			}
			if nextToken := tokens[i+1]; nextToken.Type == assign {

			}
		case value:
		}
	}

	return cmd, nil
}

func (a *App) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args)
	cmd, err := a.parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	//cmdName := args[1]
	//cmd, ok := a.commands[cmdName]
	//if !ok {
	//	return fmt.Errorf("unknown identifier")
	//}

	return cmd.Action(Context{})
}

type FlagType int

const (
	IntFlag FlagType = iota
	BoolFlag
	StringFlag
	FloatFlag
)

type Flag struct {
	Type        FlagType
	Name        string
	Description string
	Value       string
}

type Command struct {
	Name        string
	Usage       string
	Description string
	Flags       []Flag
	Action      func(ctx Context) error

	flags map[string]Flag
}

type Context struct {
	params []string
	flags  []string
}

func (Context) StringFlag(name string) {

}
