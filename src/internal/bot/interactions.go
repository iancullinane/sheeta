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

// ProcessInteraction is for any kind of interaction to get wrapped and sent
// back to match the ApiGatewayV2Prozy response format, pass in session and
// config in case they are needed
// todo::Pull out into more complex something?
func (b *bot) ProcessInteraction(interaction discordgo.Interaction) (string, error) {

	// var callback discordgo.InteractionResponse
	var resp string
	data := interaction.ApplicationCommandData()
	switch data.Name {
	case "zomboid":
		resp = b.ZomboidActionHandler(data)
	default:
		resp = "Command not found"
	}

	return resp, nil
}

func (b *bot) ZomboidActionHandler(data discordgo.ApplicationCommandInteractionData) string {
	log.Println("Zomboid Action Handler")

	// stsCreds := credentials.NewCredentials(&stscreds.AssumeRoleProvider{
	// 	Client:       sts.New(b.r.Session),
	// 	RoleARN:      roleArn,
	// 	Duration:     stscreds.DefaultDuration,
	// 	ExpiryWindow: time.Duration(float32(stscreds.DefaultDuration) * .9),
	// })

	// b.r.Session.Config.Credentials = stsCreds

	// b.r.AwsConfig.cre

	// r53 := route53.New(b.r.Session, b.r.AwsConfig)
	// hz, err := r53.ListHostedZonesByName(&route53.ListHostedZonesByNameInput{
	// 	DNSName: aws.String("adventurebrave.com"),
	// })
	// if err != nil {
	// 	if aerr, ok := err.(awserr.Error); ok {
	// 		return fmt.Sprintf("AWS error in handler: %s", aerr)
	// 	}
	// 	return fmt.Sprintf("Not AWS error: %s", err)
	// }

	return "Hello"
}
