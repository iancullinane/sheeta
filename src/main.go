package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/src/application"
	"github.com/iancullinane/sheeta/src/internal/bot"
	"github.com/iancullinane/sheeta/src/internal/chat"
	"github.com/iancullinane/sheeta/src/internal/discord"
	"github.com/iancullinane/sheeta/src/internal/exp"
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

// Use the init to provide certain values before the handler, in particular
// provide the session for this invocation
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&RunSlashBuilder, "b", "", "Slash command builder")
	flag.Parse()

	sess = session.Must(session.NewSession())
	awsCfg = &aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		Region:                        aws.String("us-east-2"), // us-east-2 is the destination bucket region
	}

	// TODO::Use motro to automate AWS keys
	ssmStore := ssm.New(sess, awsCfg)
	pKey, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/publicKey"))
	if err != nil {
		panic(err)
	}

	publicKey = *pKey.Parameter.Value
}

func Sheeta(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	if req.RawPath == "/v1/test" {
		return makeResponse(req), nil
	}

	dsess, _ := discordgo.New(publicKey)
	log.Println(dsess.LogLevel)

	validateResp, err := discord.Validate(publicKey, req)
	if validateResp != nil || err != nil {
		return *validateResp, err
	}

	var interaction discordgo.Interaction
	err = json.Unmarshal([]byte(req.Body), &interaction)
	if err != nil {
		log.Printf("interaction unmarshall: %s", err)
	}

	// All bots must be able to handle ping and validate
	if interaction.Type == discordgo.InteractionPing {
		return chat.MakePing(), nil
	}

	// Create clients to be used by modules
	// cfnClient := cloudformation.New(sess)
	// ec2Client := ec2.New(sess)
	// s3Client := s3manager.NewDownloader(sess)

	// gitClient :=
	// import "github.com/google/go-github/v44/github"

	// Instantiate modules
	availableModules := map[string]bot.Module{
		// "deploy": deploy.New(cfnClient, s3Client),
		// "server": server.New(ec2Client),
		"ctest": exp.New(),
	}

	appConfig := map[string]string{}
	bot := bot.NewBot(availableModules, sess, awsCfg, appConfig)

	var resp events.APIGatewayV2HTTPResponse

	// 	go func() {
	// 		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// 		defer cancel()
	// 		s.slackController.SubmitEvent(ctx, &outerEvent.InnerEvent, req.ApiAppId)
	// 	}()
	// 	return &slackproto.EventResponse{}, nil
	// }

	// body, err := bot.ProcessInteraction(interaction)
	// if err != nil {
	// headerSetter := make(map[string]string)
	// headerSetter["Content-Type"] = "application/json"
	// resp.StatusCode = 200
	// resp.Headers = headerSetter
	// text := fmt.Sprintf("Failed to process interaction; %v", err.Error())
	// resp.Body = string(bot.MakeResponseChannelMessageWithSource(text))
	// }

	headerSetter := make(map[string]string)
	headerSetter["Content-Type"] = "application/json"
	resp.StatusCode = 200
	resp.Headers = headerSetter

	// resp.Body = string(bot.MakeResponseChannelMessageWithSource(body))
	resp.Body = string(bot.MakeDeferredChannelMsg())

	// go func() {
	// 	body, err := bot.ProcessInteraction(interaction)
	// }()

	return resp, nil
}

func makeResponse(evt events.APIGatewayV2HTTPRequest) events.APIGatewayV2HTTPResponse {

	var resp events.APIGatewayV2HTTPResponse
	headerSetter := make(map[string]string)
	headerSetter["Content-Type"] = "application/json"
	resp.StatusCode = 200
	resp.Headers = headerSetter

	resp.Body = prettyPrint(evt)

	return resp
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
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

	if RunSlashBuilder == "delete" {
		ssmStore := ssm.New(sess, awsCfg)
		err := application.DeleteSlashCommands(ssmStore)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	lambda.Start(Sheeta)
}
