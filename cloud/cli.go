package cloud

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/bot"
	"github.com/urfave/cli/v2"
)

var (
	stackNameFlag = cli.StringFlag{
		Name:     "stack-name",
		Usage:    "Name to give the stack",
		Required: true,
	}

	templateFlag = cli.StringFlag{
		Name:     "template-name",
		Usage:    "Repository of the service",
		Required: true,
	}
)

func (cm *cloud) GenerateCLI() {

	handlers := make(map[string]bot.Handler)

	handlers["deploy"] = bot.Handler{
		DiscordFn: cm.DeployHandler,
		Flags: []cli.Flag{
			&templateFlag,
			&stackNameFlag,
		},
		ApiFn: func(c *cli.Context) error {
			err := cm.Deploy(cm.r, c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	handlers["update"] = bot.Handler{
		DiscordFn: cm.UpdateHandler,
		Flags: []cli.Flag{
			&templateFlag,
		},
	}

	handlerFuncs := make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate))
	for k, hand := range handlers {
		handlerFuncs[k] = hand.DiscordFn
	}

	newCLI := cm.buildCLI(handlers)
	cm.cliapp = &newCLI
}

func (cm *cloud) buildCLI(handlers map[string]bot.Handler) cli.App {

	var tmpCLI cli.App
	var cmds []*cli.Command
	for k, hand := range handlers {
		cmds = append(cmds, cm.buildCmd(k, "", hand))
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

func (cm *cloud) buildCmd(name, usage string, handler bot.Handler) *cli.Command {

	c := cli.Command{
		Name:   name,
		Usage:  usage,
		Flags:  handler.Flags,
		Action: handler.ApiFn,
	}

	return &c
}
