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
					Name:        "env-config",
					Description: "The config file",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "template",
					Description: "The cfn template",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "env",
					Description: "Environment to deploy into",
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
