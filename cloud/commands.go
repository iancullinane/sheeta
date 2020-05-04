package cloud

import (
	"context"
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

// ParseMessage reads the incoming mention of Sheeta and executes work
func ParseMessage(ctx context.Context, m string, args []string) error {

	// var var1 string
	// var var2 string

	messenger := &cli.App{
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "create",
				Usage: "Execute the given CloudFormation on an environment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "template-name",
						Usage: "Repository of the service",
					},
				},
				Action: func(c *cli.Context) error {

					log.Printf("Deploy: %s", m)

					return nil
				},
			},
		},

		// TODO::Put default in its own function
		Action: func(c *cli.Context) error {
			return fmt.Errorf("Not a valid command")
		},
	}

	err := messenger.RunContext(ctx, args)
	if err != nil {
		return cli.Exit(err, 86)
	}

	return nil
}
