package bot

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bwmarrin/discordgo"
)

// func MakeReturn(r discordgo.InteractionResponse, status int)

func MakeResponse(msg string) string {
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

func MakePing() events.APIGatewayV2HTTPResponse {
	callback := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseType(discordgo.InteractionPing),
		Data: nil,
	}
	responseData, err := json.Marshal(callback)
	if err != nil {
		log.Println(err)
	}

	// resp.Body = string(responseData)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(responseData),
	}
}

func ProcessInteraction(interaction discordgo.Interaction) events.APIGatewayV2HTTPResponse {
	var resp events.APIGatewayV2HTTPResponse

	// var callback discordgo.InteractionResponse
	if interaction.Type == discordgo.InteractionPing {
		return MakePing()
	}

	headerSetter := make(map[string]string)
	headerSetter["Content-Type"] = "application/json"
	resp.StatusCode = 200
	resp.Headers = headerSetter

	resp.Body = string(MakeResponse(fmt.Sprintf("Heard %s", interaction.Member.User.Username)))

	return resp
}
