# CLI

A lightweight package for building CLI apps, originally implemented for fun but feel free to use it.
Its usage is inspired by [urfave](https://github.com/urfave/cli) and [cobra](https://github.com/spf13/cobra).

### Features

* Minimal, testable, and simple code
* Commands with subcommands: `cmd [nested-sub-commands]`
* Specify flags by `--flag=value` or `--flag value`
* Specify obligatory flags and default values

### Examples

Here are a few examples of the types of commands you can create:

* `./app add 10 5`
* `./app calculate multiply 7 3`
* `./app greet --name user --upper-case true`
* `./app [greet|goodby|echo] ...`


For concrete examples checkout command_test.go.