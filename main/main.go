package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	app := Command{
		Name:        "my-cli",
		Usage:       "my-cli farewell|greet",
		Description: "prints",
		Commands:    map[string]*Command{},
	}
	app.AddCommand(&Command{
		Name:        "farewell",
		Description: "Says goodbye",
		Flags: []Flag{
			BoolFlag{
				Name:        "NICE",
				Description: "My desc",
				Required:    true,
				Default:     false,
			},
			StringFlag{
				Name:        "STR",
				Description: "Some STR addition",
				Required:    false},
		},
		Action: func(ctx Context) error {
			addition := ctx.String("STR")
			if ctx.Bool("NICE") {
				fmt.Println("You are my farewell!", addition)
			} else {
				fmt.Println("Goodbye!", addition)
			}
			return nil
		},
	})
	app.AddCommand(&Command{
		Name:        "greet",
		Usage:       "greet --name NAME",
		Description: "Outputs a greeting",
		Commands:    nil,
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
		Action: func(ctx Context) error {
			name := ctx.String("name")
			greeting := fmt.Sprintf("Hello, %s!", name)
			if shouldShout := ctx.Bool("shout"); shouldShout {
				greeting = strings.ToUpper(greeting)
			}
			fmt.Println(greeting)
			return nil
		},
	})

	fmt.Println(os.Args)
	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
