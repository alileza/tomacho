package main

import (
	"os"
	"tomato/command"
	"tomato/log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tomato",
		Usage: "Tomato functional testing tools",
		Commands: []*cli.Command{
			command.InitCommand,
			command.RunCommand,
		},
	}
	app.Version = "0.0.1"
	if err := app.Run(os.Args); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
