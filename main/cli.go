package main

import (
	"context"
	"fmt"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Commands    map[string]*Command
	Flags       []Flag
	Action      func(ctx Context) error
}

func (a *Command) AddCommand(command *Command) {
	a.Commands[command.Name] = command
}

func (a *Command) build(tokens []Token) (*Command, context.Context, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("parsing failed: missing identifier")
	}
	if tokens[0].Type != identifierType {
		return nil, nil, fmt.Errorf("parsing failed: first argument must be a identifier")
	}
	cmdName := tokens[0].Value
	cmd, ok := a.Commands[cmdName]
	if !ok {
		return nil, nil, fmt.Errorf("parsing failed: unknown identifier: %s", cmdName)
	}
	flagCtx, err := parse(a.Commands, cmd.Flags, tokens)
	return cmd, flagCtx, err
}

func (a *Command) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args[1:])
	cmd, flags, err := a.build(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	return cmd.Action(Context{
		Params: nil,
		Flags:  flags,
	})
}

type Context struct {
	Params []string
	Flags  context.Context
}

func (c Context) String(name string) string {
	return c.Flags.Value(name).(Flag).Value().(string)
}

func (c Context) Bool(name string) bool {
	return c.Flags.Value(name).(Flag).Value().(bool)
}

func (c Context) Int(name string) (value int, ok bool) {
	//data, ok := c.Flags[name]
	//if !ok {
	//	return 0, false
	//}
	//value, err := strconv.Atoi(data)
	//if err != nil {
	//	return 0, false
	//}
	return value, true
}

func (c Context) Float(name string) (value float64, ok bool) {
	//data, ok := c.Flags[name]
	//if !ok {
	//	return 0, false
	//}
	//value, err := strconv.ParseFloat(data, 32)
	//if err != nil {
	//	return 0, false
	//}
	return value, true
}
