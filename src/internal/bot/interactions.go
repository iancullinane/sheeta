package bot

import (
	"encoding/json"
	"log"

	"github.com/bwmarrin/discordgo"
)

// func MakeReturn(r discordgo.InteractionResponse, status int)

// MakeResponse is a wrapper to create a generic message back to the user
func (b *bot) MakeResponseChannelMessageWithSource(msg string) string {
	callback := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}
	responseData, err := json.Marshal(callback)
	if err != nil {
		log.Println(err)
	}
	return string(responseData)
}

func (b *bot) MakeDeferredChannelMsg() string {
	callback := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		// Data: &discordgo.InteractionResponseData{
		// 	Content: msg,
		// },
	}
	responseData, err := json.Marshal(callback)
	if err != nil {
		log.Println(err)
	}
	return string(responseData)
}

// ProcessInteraction is for any kind of interaction to get wrapped and sent
// back to match the ApiGatewayV2Prozy response format, pass in session and
// config in case they are needed
// todo::Pull out into more complex something?
func (b *bot) ProcessInteraction(interaction *discordgo.Interaction) (string, error) {

	// var callback discordgo.InteractionResponse
	var resp string
	cmd := interaction.ApplicationCommandData()

	if mod, ok := b.r.Modules[cmd.Name]; ok {
		mod.Handler(interaction, b.r.DiscrodSession)
	} else {
		resp = "No module found"
	}
	return resp, nil
}
