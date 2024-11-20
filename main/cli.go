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

func (c *Command) AddCommand(command *Command) {
	c.Commands[command.Name] = command
}

func (c *Command) prepare(token Token) (*Command, error) {
	if token.Type != identifierType {
		return nil, fmt.Errorf("parsing failed: first argument must be a identifier")
	}
	cmdName := token.Value
	cmd, ok := c.Commands[cmdName]
	if !ok {
		return nil, fmt.Errorf("parsing failed: unknown identifier: %s", cmdName)
	}
	return cmd, nil
}

func (c *Command) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args[1:])
	if len(tokens) == 0 {
		return fmt.Errorf("parsing failed: missing identifier")
	}

	cmd, err := c.prepare(tokens[0])
	if err != nil {
		return fmt.Errorf("failed to Parse input: %w", err)
	}

	p := initParser(cmd.Flags)
	flags, params, err := p.Parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse tokens: %w", err)
	}

	return cmd.Action(Context{
		Params: params,
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
