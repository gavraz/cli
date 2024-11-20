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
		flags[flag.ID()] = flag
	}

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
			if lastToken && flag.Type() != BoolFlagType {
				return nil, fmt.Errorf("parsing failed: missing argument for flag")
			}
			if lastToken {
				// just a boolean flag that is true
				continue
			}
			if nextToken := tokens[i+1]; nextToken.Type == assign {

			}
		case value:
		default:
			panic(fmt.Errorf("parsing failed: unknown token type: %s", token))
		}
	}

	return cmd, nil
}

func (a *App) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args[1:])
	cmd, err := a.parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	return cmd.Action(Context{})
}

type FlagType int

const (
	IntFlagType FlagType = iota
	BoolFlagType
	StringFlagType
	FloatFlagType
)

type Flag interface {
	ID() string
	Type() FlagType
	IsSet() bool
}

type BoolFlag struct {
	Name        string
	Description string
	Required    bool
	Default     bool

	isSet bool
	value bool
}

func (f *BoolFlag) ID() string {
	return f.Name
}

func (*BoolFlag) Type() FlagType {
	return BoolFlagType
}

func (f *BoolFlag) IsSet() bool {
	return f.isSet
}

func (f *BoolFlag) Value() bool {
	if f.isSet {
		return f.value
	}

	return f.Default
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
	Params []string
	Flags  map[string]Flag
}

func (c Context) String(name string) (value string, ok bool) {
	//value, ok = c.Flags[name]
	return
}
func (c Context) Bool(name string) (value bool) {
	f, ok := c.Flags[name]
	if !ok {
		panic(fmt.Sprintf("flag '%s' not found", name))
	}

	boolFlag, ok := f.(*BoolFlag)
	if !ok {
		panic(fmt.Sprintf("flag '%s' is not a BoolFlag", name))
	}

	return boolFlag.Value()
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
