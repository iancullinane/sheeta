package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/iancullinane/sheeta/src/internal/services"
)

// Variables used for command line parameters
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
	pKey, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/publicKey"))
	if err != nil {
		panic(err)
	}

	publicKey = *pKey.Parameter.Value
}

// func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

// 	log.Println(req.Body)
// 	log.Println(json.Marshal(req.Body))

// 	validateResp, err := discord.Validate(publicKey, req)
// 	if validateResp != nil || err != nil {
// 		return *validateResp, err
// 	}

// 	var interaction discordgo.Interaction
// 	err = json.Unmarshal([]byte(req.Body), &interaction)
// 	if err != nil {
// 		log.Printf("error: %s", err)
// 	}

// 	processedInteraction := bot.ProcessInteraction(interaction)

// 	return processedInteraction, nil
// }

type Alpha struct {
	One     string `json:"one"`
	PropTwo int    `json:"prop_2"`
}

func mainTwo() {
	testA := Alpha{One: "PropOne", PropTwo: 2}
	b, err := json.Marshal(testA)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println(string(b))

}

//
// Main
//
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// if err := json.NewEncoder(w).Encode(resp); err != nil {
		// 	http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		// 	return
		// }
		response := map[string]string{"number": "five"}
		w.Header().Set("content-type", "application/json") // and this
		json.NewEncoder(w).Encode(response)
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}

// https://stackoverflow.com/questions/52782057/golang-aws-api-gateway-invalid-character-e-looking-for-beginning-of-value
// func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

// 	resp := discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: "Some messafe",
// 		},
// 	}

// 	b, err := json.Marshal(resp)
// 	if err != nil {
// 		fmt.Println(err)
// 		// return
// 	}

// 	log.Println("passed")
// 	log.Println(string(b))

// 	return events.APIGatewayProxyResponse{
// 		Body: string(b),
// 	}, nil
// }
