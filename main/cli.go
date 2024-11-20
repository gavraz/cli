package main

import (
	"fmt"
)

type App struct {
	Name        string
	Description string
	Version     string

	commands map[string]*Command
}

func NewApp(name, description, version string) *App {
	return &App{
		Name:        name,
		Description: description,
		Version:     version,
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
			err := flags[flag.Value].Parse(value.Value)
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
			f, ok := flag.(*BoolFlag)
			if ok {
				f.Set(true)
				i += 1
				continue
			}
			lastToken := i == len(tokens)-1
			if lastToken {
				return nil, nil, fmt.Errorf("parsing failed: missing argument for flag")
			}
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

func (c Context) String(name string) string {
	f, ok := c.Flags[name]
	if !ok {
		panic(fmt.Sprintf("flag '%s' not found", name))
	}

	strFlag, ok := f.(*StringFlag)
	if !ok {
		panic(fmt.Sprintf("flag '%s' is not a string flag", name))
	}

	return strFlag.Value()
}
func (c Context) Bool(name string) bool {
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
