package cloud

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	bucketNameKey = "bucketName"
)

// ExportHandlers collects the available commands for a module for a CLI
// to consume
func (cm *cloud) ExportHandlers() []func(s *discordgo.Session, m *discordgo.MessageCreate) {
	var h []func(s *discordgo.Session, m *discordgo.MessageCreate)
	h = append(h, cm.DeployHandler)
	h = append(h, cm.UpdateHandler)
	return h
}

// DeployHandler is a handler function for the 'deploy' command
func (cm *cloud) DeployHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Just ignore certain cases like the bot mentioning itself
	if !validateMsg(m.Author.ID, s.State.User.ID, m.Mentions) {
		return
	}

	msg := strings.Split(m.ContentWithMentionsReplaced(), " ")[1:]
	if msg[1] != "deploy" {
		return
	}

	// TODO::I think think there is a better way to leverage the run function
	// of the cli library, but right now I think it is fine to use it for
	// input validation only
	err := cm.cliapp.Run(msg)
	if err != nil {
		SendErrorToUser(s, err, m.ChannelID)
		return
	}

	// input := s3.GetObjectInput{
	// 	Bucket: aws.String(cm.cfg[bucketNameKey]),
	// }

	// obj, err := cm.r.S3.GetObject(&input)
	// if err != nil {
	// 	SendErrorToUser(s, err, m.ChannelID)
	// 	return
	// }

	// reply := fmt.Sprintf("Heard! (%s) for object: %s", m.Message.Content, *obj.StorageClass)
	// s.ChannelMessageSend(m.ChannelID, reply)

	// Create(c.r)
	log.Print("Completed deploy handler")
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

	msg := strings.Split(m.ContentWithMentionsReplaced(), " ")[1:]
	if msg[1] != "update" {
		return
	}

	// TODO::I think think there is a better way to leverage the run function
	// of the cli library, but right now I think it is fine to use it for
	// input validation only
	err := cm.cliapp.Run(msg)
	if err != nil {
		SendErrorToUser(s, err, m.ChannelID)
		return
	}

	log.Print("Completed update handler")
}

// EmbedErrorMsg sends an embedded message with a red styled edge
func EmbedErrorMsg(s string) discordgo.MessageEmbed {
	var me discordgo.MessageEmbed
	me.Color = 14813706
	me.Description = s
	return me
}

// SendErrorToUser sends the contents of an error back to the user as an
// embedded message
// TODO::Send to the user privately (does discord have ephemeral messages?)
func SendErrorToUser(s *discordgo.Session, err error, channelID string) {
	errEmbed := EmbedErrorMsg(err.Error())
	msgSend := discordgo.MessageSend{
		Embed: &errEmbed,
	}
	s.ChannelMessageSendComplex(channelID, &msgSend)
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
