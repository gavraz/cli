package main

import (
	"fmt"
	"strconv"
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

func (a *App) parse(tokens []Token) (*Command, map[string]Flag, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("parsing failed: missing identifier")
	}
	if tokens[0].Type != identifierType {
		return nil, nil, fmt.Errorf("parsing failed: first argument must be a identifier")
	}
	cmdName := tokens[0].Value
	cmd, ok := a.commands[cmdName]
	if !ok {
		return nil, nil, fmt.Errorf("parsing failed: unknown identifier: %s", cmdName)
	}

	flags := make(map[string]Flag, len(cmd.Flags))
	for _, flag := range cmd.Flags {
		flags[flag.ID()] = flag
	}

	for i := 0; i < len(tokens); {
		token := tokens[i]
		switch token.Type {
		case assignType:
			if i+2 >= len(tokens) {
				return nil, nil, fmt.Errorf("parsing failed: missing argument for assign op")
			}
			flag := tokens[i+1]
			if flag.Type != flagType {
				return nil, nil, fmt.Errorf("parsing failed: unexpected operand %s: expected flag type", flag)
			}
			value := tokens[i+2]
			if value.Type != valueType {
				return nil, nil, fmt.Errorf("parsing failed: expected value type")
			}
			err := flags[flag.Value].Set(value.Value)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing failed: %w", err)
			}
			i += 3
		case identifierType:
			if _, ok := a.commands[token.Value]; !ok {
				return nil, nil, fmt.Errorf("parsing failed: unknown identifier: %s", token.Value)
			}
			i += 1
		case flagType:
			flag, ok := flags[token.Value]
			if !ok {
				return nil, nil, fmt.Errorf("parsing failed: unknown flag: %s", token.Value)
			}
			lastToken := i == len(tokens)-1
			if lastToken && flag.Type() != BoolFlagType {
				return nil, nil, fmt.Errorf("parsing failed: missing argument for flag")
			}
			boolFlag := flag.(*BoolFlag)
			_ = boolFlag.Set("true")
			i += 1
		case valueType:
			i += 1
		default:
			panic(fmt.Errorf("parsing failed: unknown token type: %s", token))
		}
	}

	return cmd, flags, nil
}

func (a *App) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments") // TODO
	}

	tokens := tokenize(args[1:])
	cmd, flags, err := a.parse(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse input: %w", err)
	}

	return cmd.Action(Context{
		Params: nil,
		Flags:  flags,
	})
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
	Set(value string) error
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

func (f *BoolFlag) Set(v string) error {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	f.value = b
	f.isSet = true
	return nil
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
