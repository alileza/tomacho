package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"tomato/config"
	"tomato/feature"
	"tomato/log"
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
				resources[r.Name] = httpclient.NewHTTPClient(r.Options)
				if err := resources[r.Name].Status(); err != nil {
					return fmt.Errorf("resource %s failed, %v", r.Name, r)
				}
				log.Info(r.Name, " initiated")
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
					if r, exist := resources[st.Resource]; exist {
						r.Exec(c.Context, st.Action, st.Arguments)
					} else {
						err = fmt.Errorf("unregistered resource of: %s", st.Resource)
					}

					log.PrintStep(st, err)
					if err != nil {
						return fmt.Errorf("execution stopped due failed step")
					}
				}
			}
			return nil
		}
		return nil
	},
}
