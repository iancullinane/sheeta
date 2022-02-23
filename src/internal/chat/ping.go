package chat

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bwmarrin/discordgo"
)

// MakePing responds to a discord ping event
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
