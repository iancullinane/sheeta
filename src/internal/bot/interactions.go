package bot

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bwmarrin/discordgo"
)

var defaultHeader map[string]string

func init() {
	defaultHeader := make(map[string]string)
	defaultHeader["Content-Type"] = "application/json"
}

func MakeResponse(msg string) string {
	callback := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}

	// Turn to string before sending back to apigateway
	// todo::figure out why proxy wasn't working and switch
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
		Headers:    defaultHeader,
		Body:       string(responseData),
	}
}

func ProcessInteraction(interaction discordgo.Interaction) events.APIGatewayV2HTTPResponse {

	if interaction.Type == discordgo.InteractionPing {
		return MakePing()
	}

	return events.APIGatewayV2HTTPResponse{
		Body:       string(MakeResponse(fmt.Sprintf("Heard %s", interaction.Member.User.Username))),
		Headers:    defaultHeader,
		StatusCode: 200,
	}
}
