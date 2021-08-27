package cloud

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/src/internal/bot"
	"github.com/urfave/cli/v2"
)

const (
	moduleName = "personality"
)

// Declare all possible flags and compose them in ExportCommands
// var (
// 	stackNameFlag = cli.StringFlag{
// 		Name:     "env",
// 		Usage:    "ENV to deploy into",
// 		Required: true,
// 	}
// )

func (self *prisoner) ExportHandler() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return self.Handler
}

func (self *prisoner) ExportCommands() []bot.Action {

	// A command has a APIFn and a DiscordFn. The API function is for doing
	// "real" work and the DiscordFn is for communication
	var r []bot.Action

	r = append(r, bot.Action{
		Name:      "play",
		Flags:     []cli.Flag{},
		DiscordFn: self.Handler,
		APIFn: func(c *cli.Context) error {
			return nil
		},
	})

	return r
}

func (self *prisoner) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Just ignore certain cases like the bot mentioning itself
	if !bot.ValidateMsg(m.Author.ID, s.State.User.ID, m.Mentions) {
		return
	}

	msg := strings.Split(m.ContentWithMentionsReplaced(), " ")[1:]
	if msg[0] != moduleName {
		bot.SendErrorToUser(s, errors.New("Invalid command"), m.ChannelID, "CLI error")
		return
	}

	if msg[1] == "play" {
		log.Println("Thing happened")
		self.playHandler(msg, s, m)
	}

}

// DeployHandler is a handler function for the 'deploy' command
func (self *prisoner) playHandler(msg []string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// This will call the ApiFn method attached to the deploy string
	err := self.cliapp.Run(msg)
	if err != nil {
		bot.SendErrorToUser(s, err, m.ChannelID, "Deploy error")
		return
	}
	bot.SendSuccessToUser(s, m.ChannelID, "Heard cap'n")
}
