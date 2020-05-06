package cloud

import (
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

	log.Printf("%#v", handlers)

	handlerFuncs := make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate))
	for k, hand := range handlers {
		handlerFuncs[k] = hand.Fn
	}
	log.Printf("%#v", handlerFuncs)
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

	return tmpCLI
}

func buildCmd(name, usage string, handler bot.Handler) *cli.Command {

	c := cli.Command{
		Name:  name,
		Usage: usage,
		Flags: handler.Flags,
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
