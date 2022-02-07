package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// type DiscordEvent struct {
// 	e discordgo.
// }
func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	// var me discordgo.MessageEmbed
	log.Println("--- Got interaction request")
	log.Printf("%#v", req.Headers)
	log.Printf("%#v", req.Body)
	publicKey := []byte("cfa20ac201afc5a130d4b5d8eabcfa186a2fe6eb6f0cc674f767a1253ec6fc63")
	typedKey, _ := hex.DecodeString("cfa20ac201afc5a130d4b5d8eabcfa186a2fe6eb6f0cc674f767a1253ec6fc63")
	log.Printf("%s", publicKey)
	var resp events.APIGatewayV2HTTPResponse
	//
	//
	//req.Headers["x-signature-ed25519"]
	signature := req.Headers["x-signature-ed25519"]
	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != ed25519.SignatureSize {
		resp.StatusCode = 401
		resp.Body = "Failed manual len check"
		return resp, err
	}

	// key = ed25519.
	// key, err := hex.DecodeString("cfa20ac201afc5a130d4b5d8eabcfa186a2fe6eb6f0cc674f767a1253ec6fc63")
	// if err != nil {
	// 	log.Println("Eror decoding")
	// 	resp.StatusCode = 401
	// 	return resp, nil
	// }

	timestamp := req.Headers["x-signature-timestamp"]
	if timestamp == "" {
		resp.StatusCode = 401
		resp.Body = "Failed on find timestamp"
		return resp, nil
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(req.Body)
	log.Println("----Write merged body")
	log.Printf("%s", msg.String())
	if !ed25519.Verify(typedKey, msg.Bytes(), sig) {
		log.Println("Should return 401 here")
		log.Printf("%v\n%v\n%v\n", publicKey, req.Headers, req.Body)
		resp.StatusCode = 401
		resp.Headers = req.Headers
		return resp, nil
	}

	log.Println("But did not return 401?")

	resp.StatusCode = 200
	resp.Body = req.Body
	return resp, nil
}

func main() {

	// d, _ := discordgo.New("Bot " + "asdfasdf")

	//
	// Mental note, make clients here, notes are below
	//

	lambda.Start(HandleRequest)
}

// // Variables used for command line parameters
// var (
// 	Token string
// )

// // For command line startup
// // TODO::Container, cloud, blah blah blah
// func init() {
// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// }

// func main() {

// 	// Set up logger to be used by package clients
// 	logger := logrus.New()
// 	logger.Level = logrus.InfoLevel
// 	logger.Out = os.Stdout

// 	// Set up config
// 	var conf *config.Config
// 	conf = conf.BuildConfigFromFile("./src/config/base.yaml")

// 	sess := session.Must(session.NewSession())
// 	// AWS config for client creation
// 	awsConfigUsEast2 := &aws.Config{
// 		CredentialsChainVerboseErrors: aws.Bool(true),
// 		S3ForcePathStyle:              aws.Bool(true),
// 		Region:                        aws.String("us-east-2"), // us-east-2 is the destination bucket region
// 	}

// 	// Create service client value configured for credentials
// 	// from assumed role.
// 	s3svc := s3manager.NewDownloader(sess)
// 	cfnSvc := cloudformation.New(sess, awsConfigUsEast2)

// 	// This effectively defines what aws services are available
// 	// TODO::I want to move this into its module, but it causes tests to break
// 	// because of a region error related to the credential chain
// 	cr := cloud.Services{
// 		S3: s3svc,
// 		CF: cfnSvc,
// 	}

// 	var bot []bot.Module
// 	c := cloud.NewCloud(cr, conf.GetValueMap())
// 	bot = append(bot, c)

// 	// Create a new Discord session using the provided bot token.
// 	d, err := discordgo.New("Bot " + Token)
// 	if err != nil {
// 		logger.Fatalf("Could not start bot: %s", err)
// 	}

// 	// Register modules handlers to discord bot
// 	for _, mod := range bot {
// 		d.AddHandler(mod.ExportHandler())
// 	}

// 	// Open a websocket connection to Discord and begin listening.
// 	err = d.Open()
// 	if err != nil {
// 		fmt.Println("error opening connection,", err)
// 		return
// 	}

// 	// Wait here until CTRL-C or other term signal is received.
// 	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
// 	<-sc

// 	// Cleanly close down the Discord session.
// 	d.Close()
// }

//
//
//
//

// logger := logrus.New()
// logger.Level = logrus.InfoLevel
// logger.Out = os.Stdout

// if config.GetVerbose() {
// 	logger.Level = logrus.DebugLevel
// }

// logger.Infof("%s %s (%s, %s)", runtimeConfig.GetEnvironment(), runtimeConfig.GetServiceName(), VersionString, runtime.Version())

// datadogAPIKey := config.GetDatadogAPIKey()
// datadogAppKey := config.GetDatadogApplicationKey()
// bb := config.GetBackupBucketName()

// sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.GetRegion())}))
// uploader := s3manager.NewUploader(sess)

// clock := clock.New()

// ddc := datadog.NewClient(datadogAPIKey, datadogAppKey)
// r := handler.Resources{
// 	DD:         ddc,
// 	Uploader:   uploader,
// 	Logger:     logger,
// 	BucketName: bb,
// 	Clock:      clock,
// }
