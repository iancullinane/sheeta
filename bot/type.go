package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

// Module is an independent set of actions containing its cli and handlers
type Module interface {
	ExportCommands() []Command
	ExportHandler() func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// Command is the function definition and its flags
type Command struct {
	Name      string
	APIFn     func(c *cli.Context) error
	DiscordFn func(s *discordgo.Session, m *discordgo.MessageCreate)
	Flags     []cli.Flag
}
