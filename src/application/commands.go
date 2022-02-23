package application

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "zomboid",
			Description: "Manage the zomboid server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "start-server",
					Description: "Start the server true/false",
					Required:    true,
				},
			},
		},
		// {
		// 	Name: "basic-command",
		// 	// All commands and options must have a description
		// 	// Commands/options without description will fail the registration
		// 	// of the command.
		// 	Description: "Basic command",
		// },
	}
)
