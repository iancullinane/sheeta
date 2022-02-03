package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
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
