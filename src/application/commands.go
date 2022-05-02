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
	}
)
