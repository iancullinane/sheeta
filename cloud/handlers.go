package cloud

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/bot"
	"github.com/urfave/cli/v2"
)

const (
	moduleName = "cloud"
)

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

func (cm *cloud) ExportCommands() []bot.Command {

	var r []bot.Command

	r = append(r, bot.Command{
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

	r = append(r, bot.Command{
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
	if !validateMsg(m.Author.ID, s.State.User.ID, m.Mentions) {
		return
	}

	msg := strings.Split(m.ContentWithMentionsReplaced(), " ")[1:]
	if msg[0] != moduleName {
		bot.SendErrorToUser(s, errors.New("Invalid command"), m.ChannelID, "CLI error")
		return
	}

	if msg[1] == "deploy" {
		log.Println("Thing happened")
		cm.DeployHandler(msg, s, m)
	}

	if msg[1] == "update" {
		cm.UpdateHandler(msg, s, m)
	}

}

// DeployHandler is a handler function for the 'deploy' command
func (cm *cloud) DeployHandler(msg []string, s *discordgo.Session, m *discordgo.MessageCreate) {

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
func (cm *cloud) UpdateHandler(msg []string, s *discordgo.Session, m *discordgo.MessageCreate) {

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

func validateMsg(authorID string, userID string, mentions []*discordgo.User) bool {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if authorID == userID {
		return false
	}

	if !containsUser(mentions, "sheeta") {
		return false
	}

	return true
}
