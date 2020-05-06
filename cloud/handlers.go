package cloud

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

const (
	bucketNameKey = "bucketName"
)

var (
	templateFlag = cli.StringFlag{
		Name:  "template-name",
		Usage: "Repository of the service",
	}
)

func (cm *cloud) GetHandlers() []func(s *discordgo.Session, m *discordgo.MessageCreate) {
	var h []func(s *discordgo.Session, m *discordgo.MessageCreate)
	h = append(h, cm.DeployHandler)
	// h = append(h, c.UpdateHandler)
	return h
}

// type Handler func(s *discordgo.Session, m *discordgo.MessageCreate)

// DeployHandler is a handler funciton for the 'deploy' command
func (cm *cloud) DeployHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !containsUser(m.Mentions, "sheeta") {
		return
	}

	log.Printf("execute for %#v", strings.Split(m.ContentWithMentionsReplaced(), " ")[1:])

	// cm.cliapp.Run(test)
	err := cm.cliapp.Run(strings.Split(m.ContentWithMentionsReplaced(), " ")[1:])
	if err != nil {
		errEmbed := PrintEmbeddedMessage(err.Error())
		msgSend := discordgo.MessageSend{
			Embed: &errEmbed,
		}
		s.ChannelMessageSendComplex(m.ChannelID, &msgSend)
		return
	}

	// input := s3.GetObjectInput{
	// 	Bucket: aws.String(c.cfg[bucketNameKey]),
	// }

	// c.r.S3.GetObject(&input)
	// tokens := strings.Split(m.ContentWithMentionsReplaced(), " ")
	// if tokens[1] == "create" {
	// 	reply := fmt.Sprintf("%s", tokens[1:])
	// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Creating stack: %s", reply))
	// 	return
	// }
	// reply := fmt.Sprintf("Heard! (%s)", m.Message.Content)
	// s.ChannelMessageSend(m.ChannelID, "Not a valid command")

	// Create(c.r)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (cm *cloud) UpdateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !containsUser(m.Mentions, "sheeta") {
		return
	}

}

// PrintEmbeddedMessage creates the data structure for a styled message
func PrintEmbeddedMessage(s string) discordgo.MessageEmbed {

	var me discordgo.MessageEmbed

	me.Color = 14813706
	me.Description = s

	return me
}
