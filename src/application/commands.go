package application

import (
	"github.com/iancullinane/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
		"responses": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Responses to a command are very important.
			// First of all, because you need to react to the interaction
			// by sending the response in 3 seconds after receiving, otherwise
			// interaction will be considered invalid and you can no longer
			// use the interaction token and ID for responding to the user's request

			content := ""
			// As you can see, the response type names used here are pretty self-explanatory,
			// but for those who want more information see the official documentation
			switch i.ApplicationCommandData().Options[0].IntValue() {
			case int64(discordgo.InteractionResponseChannelMessageWithSource):
				content =
					"You just responded to an interaction, sent a message and showed the original one. " +
						"Congratulations!"
				content +=
					"\nAlso... you can edit your response, wait 5 seconds and this message will be changed"
			default:
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
				})
				if err != nil {
					s.FollowupMessageCreate(s.State.User.ID, i.Interaction, true, &discordgo.WebhookParams{
						Content: "Something went wrong",
					})
				}
				return
			}
		},
	}
)
