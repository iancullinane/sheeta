package application

import (
	"log"

	"github.com/iancullinane/discordgo"
)

func CreateSlashCommands(appID string, s *discordgo.Session) error {

	s.AddHandler(commandHandlers["basic-command"])
	s.AddHandler(commandHandlers["responses"])
	for _, v := range commands {
		_, err := s.ApplicationCommandCreate("703973863335264286", "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		// cmd.
	}

	// whs, _ := s.ChannelWebhooks("703965708165447734")
	// for _, v := range whs {
	// 	log.Printf("%#v", v)
	// }

	// for _, v := range commands {
	// 	_, err := s.WebhookCreate("703965708165447734", "basic-command", "")
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' webhook: %v", v.Name, err)
	// 	}
	// }

	// cmds, _ := s.ApplicationCommands("703973863335264286", "")
	// for _, v := range cmds {
	// 	log.Printf("%#v", v)

	// }

	// for _, v := range cmds {
	// 	err := s.ApplicationCommandDelete(v.ApplicationID, "", v.ID)
	// 	if err != nil {
	// 		log.Printf("%#v", err)
	// 	}
	// }

	// for _, v := range commands {
	// 	_, err := s.ApplicationCommandDelete("703973863335264286")
	// 	if err != nil {
	// 		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	// 	}
	// }

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
