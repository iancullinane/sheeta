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
	Modules    map[string]Module
}

func NewBot(modules map[string]Module, sess *session.Session, aws *aws.Config, ac map[string]string) *bot {

	// activeMods := make(map[string]string, len(modules))
	// for v := range modules {
	// 	activeMods[v] =
	// }

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
	Handler(data discordgo.ApplicationCommandInteractionData) string
}

// Action is the function definition and its flags
type Action struct {
	Name      string
	APIFn     func(c *cli.Context) error
	DiscordFn func(s *discordgo.Session, m *discordgo.MessageCreate)
	Flags     []cli.Flag
}
