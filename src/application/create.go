package application

import (
	"log"

	"github.com/iancullinane/discordgo"
)

func CreateSlashCommands(appID string, s *discordgo.Session) error {

	s.AddHandler(commandHandlers["basic-command"])
	s.AddHandler(commandHandlers["responses"])
	// for _, v := range commandHandlers {

	// }

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate("703973863335264286", "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	return nil
}

// func BasicInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	// "basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: "Hey there! Congratulations, you just executed your first slash command",
// 		},
// 	})
// 	// },
// }
