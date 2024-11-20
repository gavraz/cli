package main

import (
	"fmt"
	"testing"
)

func Test_CLI(t *testing.T) {
	app := Command{
		Name:        "cli",
		Subcommands: map[string]*Command{},
	}
	app.AddSubcommand(&Command{
		Name:        "farewell",
		Description: "Says goodbye",
		Flags: []Flag{
			BoolFlag{
				Name:     "BOOL",
				Required: true,
				Default:  false,
			},
			StringFlag{
				Name:     "STR",
				Required: false,
			},
		},
		Action: func(ctx Flags, args []string) error {
			optional := ctx.String("STR")
			if ctx.Bool("BOOL") {
				fmt.Println("bool is on", optional)
			} else {
				fmt.Println("bool is off", optional)
			}
			return nil
		},
	})

}
