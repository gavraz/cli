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
		Usage:       "my-cli greet",
		Description: "prints",
		Action: func(ctx Flags, args []string) error {
			fmt.Println("Welcome to my cli!")
			return nil
		},
		Subcommands: []*Command{
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Input:	greet --name Taylor --shout true
	// Output:	HELLO, TAYLOR!
}
