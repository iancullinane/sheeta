package bot

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

type bot struct {
	r   Resources
	ctl *Controller
}

type Resources struct {
	Session    *session.Session
	AwsConfig  *aws.Config
	AppContext map[string]string //for things pulled from ssm
	Modules    map[string]Module
}

func NewBot(modules map[string]Module, sess *session.Session, aws *aws.Config, ac map[string]string) *bot {
	return &bot{
		r: Resources{
			Session:    sess,
			AwsConfig:  aws,
			AppContext: ac,
			Modules:    modules,
		},
	}
}

// Module is an independent set of actions containing its cli and handlers
type Module interface {
	Handler(discordgo.ApplicationCommandInteractionData, Controller) string
	// Handler(discordgo.ApplicationCommandInteractionData, chan string) string
}

// Action is the function definition and its flags
type Action struct {
	Name      string
	APIFn     func(c *cli.Context) error
	DiscordFn func(s *discordgo.Session, m *discordgo.MessageCreate)
	Flags     []cli.Flag
}
