package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_CLI(t *testing.T) {
	out := &bytes.Buffer{}

	app := Command{
		Name: "my-cli",
		Action: func(ctx Flags, args []string) error {
			fmt.Fprintf(out, "Welcome to my cli!")
			return nil
		},
		Subcommands: []*Command{
			{
				Name:        "farewell",
				Description: "Says goodbye",
				Flags: []Flag{
					BoolFlag{
						Name:     "UPPER",
						Required: true,
						Default:  false,
					},
				},
				Action: func(ctx Flags, args []string) error {
					msg := fmt.Sprintf("You are my farewell!")
					if ctx.Bool("UPPER") {
						msg = strings.ToUpper(msg)
					}
					fmt.Fprintf(out, msg)
					return nil
				},
			},
			{
				Name:        "greet",
				Usage:       "greet --name NAME",
				Description: "Outputs a greeting",
				Subcommands: nil,
				Flags: []Flag{
					StringFlag{
						Name:        "name",
						Description: "Your name",
						Required:    true,
					},
				},
				Action: func(ctx Flags, args []string) error {
					name := ctx.String("name")
					greeting := fmt.Sprintf("Hello, %s!", name)
					fmt.Fprintf(out, greeting)
					return nil
				},
			},
		},
	}

	addCmd := &Command{
		Name:        "add",
		Usage:       "add A B",
		Description: "calculates A + B",
		Action: func(ctx Flags, args []string) error {
			x, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "sum is: %v", x+y)
			return nil
		},
	}
	subCmd := &Command{
		Name:        "sub",
		Usage:       "sub A B",
		Description: "calculates A - B",
		Action: func(ctx Flags, args []string) error {
			x, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "diff is: %v", x-y)
			return nil
		},
	}

	calcCmd := &Command{
		Name:        "calc",
		Usage:       "calc [add|sub] X Y",
		Description: "calculates X op Y",
		Subcommands: []*Command{addCmd, subCmd},
	}
	app.AddSubcommand(calcCmd)

	assertCommand := func(t *testing.T, input string, expected string) {
		args := strings.Split(input, " ")
		out.Reset()
		err := app.Run(args)
		assert.NoError(t, err)
		assert.Equal(t, expected, out.String())
	}

	assertCommand(t, "./cli", "Welcome to my cli!")
	assertCommand(t, "./cli farewell --UPPER false", "You are my farewell!")
	assertCommand(t, "./cli farewell --UPPER true", "YOU ARE MY FAREWELL!")
	assertCommand(t, "./cli greet --name Joe", "Hello, Joe!")
	assertCommand(t, "./cli calc add 3 7", "sum is: 10")
	assertCommand(t, "./cli calc sub 3 7", "diff is: -4")
}
