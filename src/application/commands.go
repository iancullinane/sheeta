package application

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "server",
			Description: "Manage a server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "start-server",
					Description: "Start the server true/false",
					Required:    true,
				},
			},
		},
		{
			Name:        "deploy",
			Description: "Deploy a stack",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "template",
					Description: "Name of the stack",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "env",
					Description: "Environment to deploy into",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "sha",
					Description: "The sha to deploy",
				},
			},
		},
	}
)
