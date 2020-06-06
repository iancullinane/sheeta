package cloud

import (
	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/bot"
	"github.com/urfave/cli/v2"
)

var (
	stackNameFlag = cli.StringFlag{
		Name:     "env",
		Usage:    "ENV to deploy into",
		Required: true,
	}

	templateFlag = cli.StringFlag{
		Name:     "stack",
		Usage:    "Name fo the template stack yaml",
		Required: true,
	}
)

// GenerateCLI creates a cli for this module
// TODO::Automate and extract from the handlers
func (cm *cloud) GenerateCLI() {

	handlers := make(map[string]bot.Handler)

	handlers["deploy"] = bot.Handler{
		DiscordFn: cm.Handler,
		Flags: []cli.Flag{
			&templateFlag,
			&stackNameFlag,
		},
		ApiFn: func(c *cli.Context) error {
			err := cm.Deploy(cm.s, c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	handlers["update"] = bot.Handler{
		DiscordFn: cm.Handler,
		Flags: []cli.Flag{
			&templateFlag,
			&stackNameFlag,
		},
		ApiFn: func(c *cli.Context) error {
			err := cm.Update(cm.s, c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	handlerFuncs := make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate))
	for k, hand := range handlers {
		handlerFuncs[k] = hand.DiscordFn
	}

	newCLI := cm.buildCLI(handlers)
	newCLI.HideHelp = true
	newCLI.HideHelpCommand = true

	cm.cliapp = newCLI

}

func (cm *cloud) buildCLI(handlers map[string]bot.Handler) *cli.App {

	var tmpCLI cli.App
	var cmds []*cli.Command
	for k, hand := range handlers {
		cmds = append(cmds, cm.buildCmd(k, "", hand))
	}
	tmpCLI.Commands = cmds
	tmpCLI.HideHelp = true
	tmpCLI.HideHelpCommand = true
	tmpCLI.CustomAppHelpTemplate = ""
	return &tmpCLI
}

func (cm *cloud) buildCmd(name, usage string, handler bot.Handler) *cli.Command {

	c := cli.Command{
		Name:   name,
		Usage:  usage,
		Flags:  handler.Flags,
		Action: handler.ApiFn,
	}

	return &c
}
