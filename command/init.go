package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cucumber/godog/colors"
	"github.com/markbates/pkger"
	"github.com/urfave/cli/v2"
)

var InitInputs struct {
	Source    string
	Directory string
}

var InitCommand *cli.Command = &cli.Command{
	Name:        "init",
	Description: "Initialize tomato config",
	Usage:       "Initialize tomato config",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "dirname",
			Aliases:     []string{"d"},
			Usage:       "Directory name for tomato configuration",
			Destination: &InitInputs.Directory,
			Value:       "tomato",
		},
	},
	Action: func(c *cli.Context) error {
		if err := os.MkdirAll(InitInputs.Directory, 0755); err != nil {
			return err
		}
		if err := os.MkdirAll(InitInputs.Directory+"/features", 0755); err != nil {
			return err
		}

		{
			f, err := pkger.Open("/files/tomato.yaml")
			if err != nil {
				return fmt.Errorf("Failed to open gofile: %w", err)
			}
			defer f.Close()

			b, err := ioutil.ReadAll(f)
			if err != nil {
				return fmt.Errorf("Failed to read content: %w", err)
			}

			if err := os.WriteFile(InitInputs.Directory+"/tomato.yaml", b, 0644); err != nil {
				return fmt.Errorf("Failed to write file: %w", err)
			}

			fmt.Fprintf(os.Stdout, colors.Green("%s created.\n"), InitInputs.Directory+"/tomato.yaml")
		}

		{
			f, err := pkger.Open("/files/features/example.yaml")
			if err != nil {
				return fmt.Errorf("Failed to open file: %w", err)
			}
			defer f.Close()

			b, err := ioutil.ReadAll(f)
			if err != nil {
				return fmt.Errorf("Failed to read content: %w", err)
			}

			if err := os.WriteFile(InitInputs.Directory+"/features/example.yaml", b, 0644); err != nil {
				return fmt.Errorf("Failed to write file: %w", err)
			}

			fmt.Fprintf(os.Stdout, colors.Green("%s created.\n"), InitInputs.Directory+"/features/example.yaml")
		}

		return nil
	},
}
