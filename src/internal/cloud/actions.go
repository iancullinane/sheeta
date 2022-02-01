package cloud

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/src/internal/bot"
	"github.com/urfave/cli/v2"
)

const (
	moduleName = "cloud"
)

// Declare all possible flags and compose them in ExportCommands
var (
	stackNameFlag = cli.StringFlag{
		Name:     "env",
		Usage:    "ENV to deploy into",
		Required: true,
	}

	templateFlag = cli.StringFlag{
		Name:     "stack",
		Usage:    "Name fo the template stack yaml",
		Required: true,
	}
)

func (cm *cloud) ExportHandler() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return cm.Handler
}

func (cm *cloud) ExportCommands() []bot.Action {

	// A command has a APIFn and a DiscordFn. The API function is for doing
	// "real" work and the DiscordFn is for communication
	var r []bot.Action

	r = append(r, bot.Action{
		Name: "deploy",
		Flags: []cli.Flag{
			&templateFlag,
			&stackNameFlag,
		},
		DiscordFn: cm.Handler,
		APIFn: func(c *cli.Context) error {
			err := cm.Deploy(cm.s, c)
			if err != nil {
				return err
			}
			return nil
		},
	})

	r = append(r, bot.Action{
		Name: "update",
		Flags: []cli.Flag{
			&templateFlag,
			&stackNameFlag,
		},
		DiscordFn: cm.Handler,
		APIFn: func(c *cli.Context) error {
			err := cm.Update(cm.s, c)
			if err != nil {
				return err
			}
			return nil
		},
	})

	return r
}

func (cm *cloud) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Just ignore certain cases like the bot mentioning itself
	if !bot.ValidateMsg(m.Author.ID, s.State.User.ID, m.Mentions) {
		return
	}

	msg := strings.Split(m.ContentWithMentionsReplaced(), " ")[1:]

	if len(msg) <= 1 {
		bot.SendErrorToUser(s, errors.New("no command tho?"), m.ChannelID, "CLI error")
		return
	}

	if msg[0] != moduleName {
		bot.SendErrorToUser(s, errors.New("invalid command"), m.ChannelID, "CLI error")
		return
	}

	if msg[1] == "deploy" {
		cm.deployHandler(msg, s, m)
	}

	if msg[1] == "update" {
		cm.updateHandler(msg, s, m)
	}

}

// DeployHandler is a handler function for the 'deploy' command
func (cm *cloud) deployHandler(msg []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// This will call the ApiFn method attached to the deploy string
	err := cm.cliapp.Run(msg)
	if err != nil {
		bot.SendErrorToUser(s, err, m.ChannelID, "Deploy error")
		return
	}
	bot.SendSuccessToUser(s, m.ChannelID, "Heard cap'n")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (cm *cloud) updateHandler(msg []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// TODO::I think think there is a better way to leverage the run function
	// of the cli library, but right now I think it is fine to use it for
	// input validation only
	err := cm.cliapp.Run(msg)
	if err != nil {
		bot.SendErrorToUser(s, err, m.ChannelID, "Update error")
		return
	}

	bot.SendSuccessToUser(s, m.ChannelID, "Heard cap'n")
}
