package bot

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

type bot struct {
	r Resources
}

type Resources struct {
	Session    *session.Session
	AwsConfig  *aws.Config
	AppContext map[string]string //for things pulled from ssm
}

func NewBot(sess *session.Session, aws *aws.Config, ac map[string]string) *bot {
	return &bot{
		r: Resources{
			Session:    sess,
			AwsConfig:  aws,
			AppContext: ac,
		},
	}
}

// Module is an independent set of actions containing its cli and handlers
type Module interface {
	ExportCommands() []Action
	ExportHandler() func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// Action is the function definition and its flags
type Action struct {
	Name      string
	APIFn     func(c *cli.Context) error
	DiscordFn func(s *discordgo.Session, m *discordgo.MessageCreate)
	Flags     []cli.Flag
}
