package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	app := NewApp("my-cli", "my CLI demo", "1.1")
	app.AddCommand(&Command{
		Name:        "farewell",
		Description: "Says goodbye",
		Action: func(ctx Context) error {
			if ctx.Bool("NICE") {
				fmt.Println("You are my farewell!")
			} else {
				fmt.Println("Goodbye!")
			}
			return nil
		},
		Flags: []Flag{
			&BoolFlag{
				Name:        "NICE",
				Description: "My desc",
				Required:    true,
				Default:     false,
			},
		},
	})
	//	app.AddCommand(&Command{
	//		Name:        "greet",
	//		Usage:       "greet --name NAME",
	//		Description: "Outputs a greeting",
	//		Flags: []clifave.Flag{
	//			&clifave.StringFlag{
	//				Name:     "name",
	//				Aliases:  []string{"n"},
	//				Usage:    "Your name",
	//				Required: true,
	//			},
	//			&clifave.BoolFlag{
	//				Name:  "shout",
	//				Usage: "Shout the greeting in uppercase",
	//				Value: false,
	//			},
	//		},
	//		Action: func(c Context) error {
	//			name := c.String("name")
	//			greeting := fmt.Sprintf("Hello, %s!", name)
	//			if shouldShout := c.Bool("shout") {
	//				greeting = strings.ToUpper(greeting)
	//			}
	//			fmt.Println(greeting)
	//			return nil
	//		},
	//	})
	//
	//Action:
	//	func(ctx Context) {
	//		// calculate 3 5 --op add
	//		x := ctx.Params[0]
	//		y := ctx.Params[1]
	//		// no? err
	//		if ctx.String("op") {
	//			fmt.Println("adding: ", x+y)
	//		}
	//	}

	// ./app greet --name
	// Define the CLI application
	//app := &clifave.App{
	//	Name:    "mycli",
	//	Usage:   "A simple example CLI application",
	//	Version: "1.0.0",
	//	Commands: []*clifave.Command{
	//		{
	//			Name:    "greet",
	//			Usage:   "Outputs a greeting",
	//			Aliases: []string{"g"},
	//			Flags: []clifave.Flag{
	//				&clifave.StringFlag{
	//					Name:     "name",
	//					Aliases:  []string{"n"},
	//					Usage:    "Your name",
	//					Required: true,
	//				},
	//				&clifave.BoolFlag{
	//					Name:  "shout",
	//					Usage: "Shout the greeting in uppercase",
	//					Value: false,
	//				},
	//			},
	//			Action: func(c *clifave.Context) error {
	//				name := c.String("name")
	//				greeting := fmt.Sprintf("Hello, %s!", name)
	//				if c.Bool("shout") {
	//					greeting = strings.ToUpper(greeting)
	//				}
	//				fmt.Println(greeting)
	//				return nil
	//			},
	//		},
	//		{
	//			Name:  "farewell",
	//			Usage: "Says goodbye",
	//			Action: func(c *clifave.Context) error {
	//				fmt.Println("Goodbye!")
	//				return nil
	//			},
	//		},
	//	},
	//}
	fmt.Println(os.Args)
	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
