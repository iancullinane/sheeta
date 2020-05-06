package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

// Module defines necessary components to be loaded as a bot module
type Module struct {
	Name     string
	Handlers map[string]func(s *discordgo.Session, m *discordgo.MessageCreate)
	CLI      *cli.App
}

// NewModule returns a new module for the main bot
func NewModule(
	name string,
	handlers map[string]Handler) *Module {

	handlerFuncs := make(map[string]func(s *discordgo.Session, m *discordgo.MessageCreate))
	for k, hand := range handlers {
		log.Printf("%#v %#v", k, hand)
		handlerFuncs[k] = hand.Fn
	}

	var m Module
	m.Name = name
	m.Handlers = handlerFuncs
	m.CLI = m.buildCLI(handlers)

	return &m
}

func (m *Module) buildCLI(handlers map[string]Handler) *cli.App {

	var tmpCLI cli.App

	cmds := make([]*cli.Command, len(handlers))

	for k, hand := range handlers {
		cmds = append(cmds, buildCmd(k, "", hand))
	}

	tmpCLI.Commands = cmds

	return &tmpCLI
}

func buildCmd(name, usage string, handler Handler) *cli.Command {

	c := &cli.Command{
		Name:  name,
		Usage: usage,
		Flags: handler.Flags,
		Action: func(c *cli.Context) error {
			log.Printf("Deploy")
			return nil
		},
	}

	return c
}
