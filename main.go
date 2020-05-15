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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/cloud"
	"github.com/iancullinane/sheeta/config"
	"github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token string
)

// Module is an independent set of actions containing its cli and handlers
type Module interface {
	GenerateCLI()
	ExportHandlers() []func(s *discordgo.Session, m *discordgo.MessageCreate)
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Set up logger to be used by package clients
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Out = os.Stdout

	// Set up config
	var conf *config.Config
	conf = conf.BuildConfigFromFile("./config/base.yaml")

	// Create a new Discord session using the provided bot token.
	d, err := discordgo.New("Bot " + Token)
	if err != nil {
		logger.Fatalf("Could not start bot: %s", err)
	}

	sess := session.Must(session.NewSession())

	// AWS config for client creation
	awsConfigUsEast2 := &aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		Region:                        aws.String("us-east-2"), // us-east-2 is the destination bucket region
	}
	// Create service client value configured for credentials
	// from assumed role.
	// s3svc := s3.New(sess, awsConfigUsEast2)
	s3svc := s3manager.NewDownloader(sess)
	cfnSvc := cloudformation.New(sess, awsConfigUsEast2)

	cr := cloud.Resources{
		S3:     s3svc,
		CF:     cfnSvc,
		Logger: logger,
	}

	// Any module must fit the module definition of retreiving handlers,
	// and generating a CLI
	var bot []Module
	c := cloud.NewCloud(cr, conf.GetValueMap())
	c.GenerateCLI()

	bot = append(bot, c)

	// Register the messageCreate func as a callback for MessageCreate events.
	for _, mod := range bot {
		for _, mod := range mod.ExportHandlers() {
			d.AddHandler(mod)
		}
	}

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
