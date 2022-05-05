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
	cmd := interaction.ApplicationCommandData()

	if mod, ok := b.r.Modules[cmd.Name]; ok {
		resp = mod.Handler(cmd)
	} else {
		resp = "No module found"
	}
	return resp, nil
}

// func (b *bot) ServerActionHandler(data discordgo.ApplicationCommandInteractionData) string {
// 	log.Println("Do something on a server")

// 	svc := ec2.New(b.r.Session)

// 	input := &ec2.DescribeInstancesInput{
// 		Filters: []*ec2.Filter{
// 			{
// 				Name: aws.String("tag:sheeta"),
// 				Values: []*string{
// 					aws.String("server-sorrow"),
// 				},
// 			},
// 		},
// 	}

// 	result, err := svc.DescribeInstances(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			default:
// 				fmt.Println(aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			fmt.Println(err.Error())
// 		}
// 		return err.Error()
// 	}

// 	log.Println(result)

// 	return *result.Reservations[0].Instances[0].InstanceId
// }
