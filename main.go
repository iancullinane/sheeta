package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/cloud"
	"github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Set up logger to be used by package clients
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Out = os.Stdout

	// Create a new Discord session using the provided bot token.
	d, err := discordgo.New("Bot " + Token)
	if err != nil {
		logger.Fatalf("Could not start bot: %s", err)
	}

	// AWS config for client creation
	awsConfigUsEast2 := &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("us-east-2"), // us-east-2 is the destination bucket region
	}
	stsSessionUsEast2 := session.Must(session.NewSession(awsConfigUsEast2))

	r := cloud.Resources{
		S3:     s3.New(stsSessionUsEast2),
		CF:     cloudformation.New(stsSessionUsEast2),
		Logger: logger,
	}
	c := cloud.NewCloud(r)

	// Register the messageCreate func as a callback for MessageCreate events.
	d.AddHandler(c.Handler)

	// Open a websocket connection to Discord and begin listening.
	err = d.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	d.Close()
}
