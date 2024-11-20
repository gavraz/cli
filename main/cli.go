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

func (c *Command) navigateToMostInnerCommand(tokens []Token, found bool, currCmd *Command) (int, *Command) {
	var i int
	for i = 1; i < len(tokens); i++ {
		var cmd *Command
		cmd, found = currCmd.nextSubcommand(tokens[i])
		if !found {
			break
		}
		currCmd = cmd
	}
	return i, currCmd
}

func (c *Command) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args[1:])
	if len(tokens) == 0 {
		return fmt.Errorf("parsing failed: missing identifier")
	}

	currCmd, found := c.nextSubcommand(tokens[0])
	if !found {
		return fmt.Errorf("failed to run command: command not found: %s", tokens[0])
	}

	i, currCmd := c.navigateToMostInnerCommand(tokens, found, currCmd)
	p := initParser(currCmd.Flags)
	ctx, args, err := p.Parse(tokens[i:])
	if err != nil {
		return fmt.Errorf("failed to parse tokens: %w", err)
	}

	return currCmd.Action(Flags{ctx: ctx}, args)
}
