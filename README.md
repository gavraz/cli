# CLI

A lightweight package for building CLI apps, originally implemented for fun but feel free to use it.
Its usage is inspired by [urfave](https://github.com/urfave/cli) and [cobra](https://github.com/spf13/cobra).

### Features

* Minimal, testable, and simple code
* Commands with subcommands: `cmd [nested-sub-commands]`
* Specify flags by `--flag=value` or `--flag value`
* Specify obligatory flags and default values

### Examples

#### Command Samples

Here are a few examples of the types of commands you can create:

* `./app add 10 5`
* `./app calculate multiply 7 3`
* `./app greet --name user --upper-case true`
* `./app [greet|goodby|echo] ...`

#### A Greeting Application

```go
package main

import (
	"fmt"
	"github.com/gavraz/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.Command{
		Name:  "my-cli",
		Usage: "my-cli [greet]",
		Action: func(ctx cli.Flags, args []string) error {
			fmt.Println("Welcome to my cli!")
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:        "greet",
				Usage:       "greet --name NAME",
				Description: "Outputs a greeting",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "name",
						Description: "Your name",
						Required:    true,
					},
					cli.BoolFlag{
						Name:        "shout",
						Description: "Shout the greeting in uppercase",
						Default:     false,
					},
				},
				Action: func(ctx cli.Flags, args []string) error {
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

```

For additional examples checkout command_test.go.