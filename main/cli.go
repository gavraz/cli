package main

import (
	"fmt"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Subcommands []*Command
	Params      []string
	Flags       []Flag
	Action      func(ctx Flags, args []string) error
}

func (c *Command) AddSubcommand(cmd *Command) {
	c.Subcommands = append(c.Subcommands, cmd)
}

func (c *Command) subcommand(name string) *Command {
	for _, cmd := range c.Subcommands {
		if cmd.Name == name {
			return cmd
		}
	}

	return nil
}

func (c *Command) nextSubcommand(token Token) (*Command, bool) {
	if token.Type != identifierType {
		return nil, false
	}
	cmdName := token.Value
	cmd := c.subcommand(cmdName)
	if cmd == nil {
		return nil, false
	}
	return cmd, true
}

func (c *Command) navigateToMostInnerCommand(tokens []Token) (int, *Command) {
	currCmd := c
	var i int
	for i = 1; i < len(tokens); i++ {
		var cmd *Command
		cmd, found := currCmd.nextSubcommand(tokens[i])
		if !found {
			break
		}
		currCmd = cmd
	}
	return i, currCmd
}

func (c *Command) Run(args []string) error {
	tokens := tokenize(args)
	i, currCmd := c.navigateToMostInnerCommand(tokens)
	p := initParser(currCmd.Flags)
	ctx, args, err := p.Parse(tokens[i:])
	if err != nil {
		return fmt.Errorf("failed to parse tokens: %w", err)
	}

	return currCmd.Action(Flags{ctx: ctx}, args)
}
