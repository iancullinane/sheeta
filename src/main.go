package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/iancullinane/discordgo"
	"github.com/iancullinane/sheeta/src/application"
	"github.com/iancullinane/sheeta/src/internal/discord"
	"github.com/iancullinane/sheeta/src/internal/services"
)

// // Variables used for command line parameters
var (
	Token           string
	RunSlashBuilder string
)

var sess *session.Session
var awsCfg *aws.Config
var publicKey string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&RunSlashBuilder, "b", "", "Slash command builder")
	flag.Parse()

	sess = session.Must(session.NewSession())
	awsCfg = &aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		Region:                        aws.String("us-east-1"), // us-east-2 is the destination bucket region
	}

	ssmStore := ssm.New(sess, awsCfg)
	pKey, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/token"))
	if err != nil {
		panic(err)
	}
	publicKey = *pKey.Parameter.Value
}

func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	validateResp, err := discord.Validate(publicKey, req)
	if validateResp != nil || err != nil {
		return *validateResp, err
	}

	var interaction discordgo.Interaction
	err = json.Unmarshal([]byte(req.Body), &interaction)
	if err != nil {
		log.Printf("error: %s", err)
	}
	var dResp discordgo.InteractionResponse
	var resp events.APIGatewayV2HTTPResponse
	if interaction.Type == discordgo.InteractionPing {
		dResp = discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseType(interaction.Type),
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		}
	} else {
		dResp = discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Heard %s", interaction.Member.User.Username),
			},
		}

		headerSetter := make(map[string]string)
		headerSetter["Content-Type"] = "application/json"
		resp.StatusCode = 200
		resp.Headers = headerSetter
	}

	responseData, err := json.Marshal(dResp)
	if err != nil {
		log.Println(err)
	}

	resp.Body = string(responseData)
	return resp, nil
}

//
// Main
//
func main() {

	// Alternate run command to build the webhooks and interactions in Discord
	if RunSlashBuilder == "create" {
		ssmStore := ssm.New(sess, awsCfg)
		err := application.CreateSlashCommands(ssmStore)
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}

	lambda.Start(HandleRequest)
}
