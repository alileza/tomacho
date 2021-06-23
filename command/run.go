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
	Verbosity  string
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
		&cli.StringFlag{
			Name:        "verbosity",
			Aliases:     []string{"v"},
			Usage:       "",
			Destination: &RunInputs.Verbosity,
			Value:       "info",
		},
	},
	Action: func(c *cli.Context) error {
		conf, err := config.Retrieve(RunInputs.ConfigPath)
		if err != nil {
			return fmt.Errorf("failed to retrieve config: %w", err)
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
			for _, sc := range ff.Scenarios {
				c.Context = resource.SetExecID(c.Context, sc.ID)
				for _, st := range sc.Steps {
					if err := resources[st.Resource].Exec(c.Context, st.Action, st.Arguments); err != nil {
						fmt.Printf("🔥 %s %s\n", colors.Red(st.ID), err)
						return fmt.Errorf("execution stopped due failed step")
					} else {
						dump, err := resources[st.Resource].DumpStorage()
						if err != nil {
							return fmt.Errorf("failed to dump storage: %w", err)
						}
						fmt.Printf("✅ %s (%s)\n", colors.Green(st.ID), dump)
					}
				}
			}
			return nil
		}

		return nil
	},
}
