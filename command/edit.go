package command

import (
	"github.com/urfave/cli/v2"
)

var EditInputs struct {
	ConfigPath string
}

var EditCommand *cli.Command = &cli.Command{
	Name:        "init",
	Description: "Initialize tomato config",
	Usage:       "Initialize tomato config",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Config path",
			Destination: &EditInputs.ConfigPath,
			Value:       "tomato.yaml",
		},
	},
	Action: func(c *cli.Context) error {

		return nil
	},
}
