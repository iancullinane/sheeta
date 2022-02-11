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
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        "server",
					Description: "Control the server",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "run",
							Description: "Start and start the zomboid server",
							Type:        discordgo.ApplicationCommandOptionBoolean,
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
					Name:        "scream",
					Description: "Commands related to server access",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "for-help",
							Description: "Get server status",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
						{
							Name:        "to-be-found",
							Description: "Give IP for security group",
							Type:        discordgo.ApplicationCommandOptionString,
							Required:    true,
						},
					},
				},
				// {
				// 	Type:        discordgo.ApplicationCommandOptionBoolean,
				// 	Name:        "run-server",
				// 	Description: "Start the server true/false",
				// 	Required:    true,
				// },
			},
		},
	}
)
