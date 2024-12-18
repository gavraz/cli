package cli

import (
	"fmt"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Subcommands []*Command
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
	for i = 0; i < len(tokens); i++ {
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
	if len(args) < 1 {
		return fmt.Errorf("missing arguments")
	}
	tokens := tokenize(args[1:])
	i, currCmd := c.navigateToMostInnerCommand(tokens)
	p := initParser(currCmd.Flags)
	if i > 0 {
		tokens = tokens[i:]
	}
	ctx, args, err := p.Parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse tokens: %w", err)
	}

	return currCmd.Action(Flags{ctx: ctx}, args)
}
