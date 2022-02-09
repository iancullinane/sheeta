package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"flag"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/iancullinane/discordgo"
	"github.com/iancullinane/sheeta/src/internal/services"
)

// // Variables used for command line parameters
var (
	Token           string
	RunSlashBuilder string
)

var (
	sess      *session.Session
	publicKey string
	ginLambda *ginadapter.GinLambdaV2
)

type Response struct {
	StatusCode int
}

// For command line startup
// TODO::Container, cloud, blah blah blah

func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	log.Printf("%#v", req.Body)

	typedKey, _ := hex.DecodeString("cfa20ac201afc5a130d4b5d8eabcfa186a2fe6eb6f0cc674f767a1253ec6fc63")

	var resp events.APIGatewayV2HTTPResponse

	signature := req.Headers["x-signature-ed25519"]
	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != ed25519.SignatureSize {
		resp.StatusCode = 401
		resp.Body = "Failed manual len check"
		return resp, err
	}

	timestamp := req.Headers["x-signature-timestamp"]
	if timestamp == "" {
		resp.StatusCode = 401
		resp.Body = "Failed on find timestamp"
		return resp, nil
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(req.Body)
	if !ed25519.Verify(typedKey, msg.Bytes(), sig) {
		resp.StatusCode = 401
		resp.Headers = req.Headers
		return resp, nil
	}

	resp.StatusCode = 200
	resp.Body = req.Body
	return resp, nil
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&RunSlashBuilder, "b", "", "Slash command builder")
	flag.Parse()
}

func init() {

	sess = session.Must(session.NewSession())
	awsConfigUsEast1 := &aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		Region:                        aws.String("us-east-1"), // us-east-2 is the destination bucket region
	}

	ssmStore := ssm.New(sess, awsConfigUsEast1)
	keyFromSSM, err := services.GetParameterDecrypted(ssmStore, aws.String("/discord/sheeta/publicKey"))
	if err != nil {
		log.Println("Error getting publickey")
		panic(err)
	}
	// typedKey, err := hex.DecodeString(*keyFromSSM.Parameter.Value)
	// if err != nil {
	// 	panic(err)
	// }

	log.Println("From init")
	log.Println(*keyFromSSM.Parameter.Value)

	publicKey = *keyFromSSM.Parameter.Value
	typedKey, _ := hex.DecodeString(publicKey)
	// stdout and stderr are sent to AWS CloudWatch Logs

	log.Printf("Gin cold start")
	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		if !discordgo.VerifyInteraction(c.Request, typedKey) {
			log.Println("failed verify")
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		log.Println("passed verify"
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginLambda = ginadapter.NewV2(r)

}

//
// Main
//

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {

	// decoded, err := hex.DecodeString(publicKey)
	// if err != nil {
	// 	log.Println("error decoding publicKey")
	// }
	// log.Println(decoded)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	log.Println("---- Make request")
	// 	log.Printf("%#v", r.Header.Get("X-Signature-Ed25519"))
	// 	log.Printf("%#v", r.Header.Get("X-Signature-Timestamp"))

	// 	log.Printf("%#v", r)
	// 	log.Println("-------------")

	// 	if discordgo.VerifyInteraction(r, decoded) {
	// 		log.Println("verified")
	// 		// w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		// json.NewEncoder(w).Encode(r.Body)
	// 		return
	// 	} else {
	// 		log.Println("not verified")
	// 		http.Error(w, "not verified", http.StatusUnauthorized)
	// 		return
	// 	}

	// })

	lambda.Start(Handler)

	// sess := session.Must(session.NewSession())
	// // AWS config for client creation
	// awsConfigUsEast1 := &aws.Config{
	// 	CredentialsChainVerboseErrors: aws.Bool(true),
	// 	S3ForcePathStyle:              aws.Bool(true),
	// 	Region:                        aws.String("us-east-1"), // us-east-2 is the destination bucket region
	// }

	// // Create service client value configured for credentials
	// // from assumed role.
	// // s3svc := s3manager.NewDownloader(sess)
	// // cfnSvc := cloudformation.New(sess, awsConfigUsEast2)
	// ssmStore := ssm.New(sess, awsConfigUsEast1)
	// dToken, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/token"))
	// if err != nil {
	// 	panic(err)
	// }

	// apiID, err := services.GetParameter(ssmStore, aws.String("/discord/sheeta/app-id"))
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println("token value")
	// log.Printf("%#v", *dToken.Parameter.Value)
	// d, err := discordgo.New("Bot " + *dToken.Parameter.Value)
	// if err != nil {
	// 	panic(err)
	// }

	// if RunSlashBuilder == "create" {
	// 	log.Println("api value")
	// 	log.Printf("%#v", *apiID.Parameter.Value)
	// 	err := application.CreateSlashCommands(*apiID.Parameter.Value, d)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	os.Exit(0)
	// }

	// lambda.Start(HandleRequest)
}
