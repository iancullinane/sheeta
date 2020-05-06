package cloud

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/bot"
	"github.com/urfave/cli/v2"
)

type cloud struct {
	r      Resources
	cfg    map[string]string
	cliapp *cli.App
}

// Resources are API's needed to execute a task
type Resources struct {
	S3     S3Client
	CF     CFClient
	Logger Logger
}

// NewCloud returns a new cloud client
func NewCloud(r Resources, cfg map[string]string) *cloud {

	return &cloud{
		r:   r,
		cfg: cfg,
	}
}

func (c *cloud) GenerateCLI() {

	handlers := make(map[string]bot.Handler)

	handlers["deploy"] = bot.Handler{
		Fn: c.DeployHandler,
		Flags: []cli.Flag{
			&templateFlag,
		},
	}

	// handlers["create"] = bot.Handler{
	// 	Fn: c.UpdateHandler,
	// 	Flags: []cli.Flag{
	// 		&templateFlag,
	// 	},
	// }

	handlerFuncs := make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate))
	for k, hand := range handlers {
		handlerFuncs[k] = hand.Fn
	}

	newCLI := c.buildCLI(handlers)
	c.cliapp = &newCLI
}

func (c *cloud) buildCLI(handlers map[string]bot.Handler) cli.App {

	var tmpCLI cli.App
	var cmds []*cli.Command
	for k, hand := range handlers {
		cmds = append(cmds, buildCmd(k, "", hand))
	}
	tmpCLI.Commands = cmds

	tmpCLI.Action = func(c *cli.Context) error {
		log.Println("Error parsing")
		return fmt.Errorf("Not a valid command")
	}

	tmpCLI.Name = "cloud"
	tmpCLI.HideHelp = true
	return tmpCLI
}

func buildCmd(name, usage string, handler bot.Handler) *cli.Command {

	log.Printf("%#v", handler)

	c := cli.Command{
		Name:  name,
		Usage: usage,
		Flags: handler.Flags,
		Action: func(c *cli.Context) error {
			log.Printf("This is the %s command", name)
			return nil
		},
	}

	return &c
}

// func (c *cloud) ExportModule(name string, r Resources) *bot.Module {

// 	handlers := make(map[string]bot.Handler)

// 	handlers["deploy"] = bot.Handler{
// 		Fn: c.DeployHandler,
// 		Flags: []cli.Flag{
// 			&templateFlag,
// 		},
// 	}

// 	handlers["create"] = bot.Handler{
// 		Fn: c.UpdateHandler,
// 		Flags: []cli.Flag{
// 			&templateFlag,
// 		},
// 	}

// 	cloudMod := bot.NewModule("cloud", handlers)

// 	return cloudMod
// }

// messenger := &cli.App{
// 	Commands: []*cli.Command{
// 		&cli.Command{
// 			Name:  "deploy",
// 			Usage: "Execute the given CloudFormation on an environment",
// 			Flags: []cli.Flag{
// 				&cli.StringFlag{
// 					Name:  "template-name",
// 					Usage: "Repository of the service",
// 				},
// 			},
// 			Action: func(c *cli.Context) error {

// 				log.Printf("Deploy: %s", m)

// 				return nil
// 			},
// 		},
// 	},

// 	// TODO::Put default in its own function
// 	Action: func(c *cli.Context) error {
// 		return fmt.Errorf("Not a valid command")
// 	},
// }

// err := messenger.RunContext(ctx, args)
// if err != nil {
// 	return cli.Exit(err, 86)
// }
