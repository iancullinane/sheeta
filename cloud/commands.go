package cloud

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

// func(c *cloud) AddCommand

// GetCLI returns a CLI definition, conforms to modularized version
func (c *cloud) GetCLI() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "deploy",
				Usage: "Execute the given CloudFormation on an environment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "template-name",
						Usage: "Repository of the service",
					},
				},
				Action: func(c *cli.Context) error {

					log.Printf("Deploy")

					return nil
				},
			},
		},

		// TODO::Put default in its own function
		Action: func(c *cli.Context) error {
			return fmt.Errorf("Not a valid command")
		},
	}
}
