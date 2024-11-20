package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := Command{
		Name:        "my-cli",
		Usage:       "my-cli farewell|greet",
		Description: "prints",
		Action: func(ctx Flags, args []string) error {
			fmt.Println("Welcome to my cli!")
			return nil
		},
		Subcommands: []*Command{
			{
				Name:        "farewell",
				Description: "Says goodbye",
				Flags: []Flag{
					BoolFlag{
						Name:     "NICE",
						Required: true,
						Default:  false,
					},
					StringFlag{
						Name:     "STR",
						Required: false,
					},
				},
				Action: func(ctx Flags, args []string) error {
					addition := ctx.String("STR")
					if ctx.Bool("NICE") {
						fmt.Println("You are my farewell!", addition)
					} else {
						fmt.Println("Goodbye!", addition)
					}
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
					BoolFlag{
						Name:        "shout",
						Description: "Shout the greeting in uppercase",
						Default:     false,
					},
				},
				Action: func(ctx Flags, args []string) error {
					name := ctx.String("name")
					greeting := fmt.Sprintf("Hello, %s!", name)
					if shouldShout := ctx.Bool("shout"); shouldShout {
						greeting = strings.ToUpper(greeting)
					}
					fmt.Println(greeting)
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
			fmt.Println("sum is:", x+y)
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
			fmt.Println("diff is:", x-y)
			return nil
		},
	}

	calcCmd := &Command{
		Name:        "calc",
		Usage:       "calc [add|sub] X Y",
		Description: "calculates X op Y",
		Subcommands: []*Command{addCmd, subCmd},
		Action: func(ctx Flags, args []string) error {
			// Run subcommand
			return nil
		},
	}
	//calcCmd.AddSubCommand(addCmd)
	//calcCmd.AddSubCommand(subCmd)
	app.AddSubcommand(calcCmd)

	fmt.Println(os.Args)
	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
