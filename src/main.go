package main

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bwmarrin/discordgo"
)

// type DiscordEvent struct {
// 	e discordgo.
// }

type InteractionResp struct {
	StatusCode int `'json:"statusCode"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (InteractionResp, error) {

	// var me discordgo.MessageEmbed
	log.Printf("%#v", req)

	//
	//
	//req.Headers["x-signature-ed25519"]
	signature := req.Headers["x-signature-ed25519"]
	sig, _ := hex.DecodeString(signature)
	if len(sig) != ed25519.SignatureSize {
		return InteractionResp{StatusCode: 401}, nil
	}

	key, _ := hex.DecodeString("cfa20ac201afc5a130d4b5d8eabcfa186a2fe6eb6f0cc674f767a1253ec6fc63")

	timestamp := req.Headers["x-signature-timestamp"]
	if timestamp == "" {
		return InteractionResp{StatusCode: 401}, nil
	}

	if !ed25519.Verify(key, []byte(req.Body), sig) {
		log.Println("Should return 401 here")
		return InteractionResp{StatusCode: 401}, nil
	}

	// resp.Body = req.Body
	return InteractionResp{StatusCode: 200}, nil
}

func main() {

	d, _ := discordgo.New("Bot " + "asdfasdf")
	log.Printf("%#v", d)
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
