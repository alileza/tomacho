package command

import (
	"fmt"

	"github.com/cucumber/godog/colors"
	"github.com/urfave/cli/v2"

	"tomato/config"
	"tomato/feature"
	"tomato/resource"
	"tomato/resource/httpclient"
)

var RunInputs struct {
	ConfigPath string
}

var RunCommand *cli.Command = &cli.Command{
	Name:        "run",
	Description: "Run tomato testing suite",
	Usage:       "Run tomato testing suite",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Config path",
			Destination: &RunInputs.ConfigPath,
			Value:       "tomato.yaml",
		},
	},
	Action: func(c *cli.Context) error {
		conf, err := config.Retrieve(RunInputs.ConfigPath)
		if err != nil {
			return fmt.Errorf("Failed to retrieve config: %w", err)
		}

		resources := make(map[string]resource.Resource)
		for _, r := range conf.Resources {
			switch r.Type {
			case "httpclient":
				// fmt.Printf("[%s] Initiating\n", r.Name)
				resources[r.Name] = httpclient.NewHTTPClient(r.Options)
				if err := resources[r.Name].Status(); err != nil {
					return fmt.Errorf("resource %s failed, %v", r.Name, r)
				}
			default:
				return fmt.Errorf("resource %s not found", r.Type)
			}
		}

		if f := c.Args().First(); f != "" {
			ff, err := feature.Retrieve(f)
			if err != nil {
				return err
			}
			c.Context = resource.SetExecID(c.Context, f)
			for _, sc := range ff.Scenarios {
				if err := resources[sc.Resource].Exec(c.Context, sc.Action, sc.Arguments); err != nil {
					fmt.Printf("[%s] %s\n", colors.Red(sc.ID), err)
					return fmt.Errorf("execution stopped due failed step")
				} else {
					fmt.Printf("[%s] Good\n", colors.Green(sc.ID))
				}
			}
			return nil
		}

		return nil
	},
}
