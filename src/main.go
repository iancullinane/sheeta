package main

import (
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

// // Variables used for command line parameters
// var (
// 	Token           string
// 	RunSlashBuilder string
// )

// var sess *session.Session
// var awsCfg *aws.Config
// var publicKey string

// func init() {
// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.StringVar(&RunSlashBuilder, "b", "", "Slash command builder")
// 	flag.Parse()

// 	sess = session.Must(session.NewSession())
// 	awsCfg = &aws.Config{
// 		CredentialsChainVerboseErrors: aws.Bool(true),
// 		S3ForcePathStyle:              aws.Bool(true),
// 		Region:                        aws.String("us-east-1"), // us-east-2 is the destination bucket region
// 	}

// 	ssmStore := ssm.New(sess, awsCfg)
// 	pKey, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/publicKey"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	publicKey = *pKey.Parameter.Value
// }

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

//
// Main
//
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)

	// // Alternate run command to build the webhooks and interactions in Discord
	// if RunSlashBuilder == "create" {
	// 	ssmStore := ssm.New(sess, awsCfg)
	// 	err := application.CreateSlashCommands(ssmStore)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	os.Exit(0)
	// }

	// lambda.Start(HandleRequest)
}
