package cloud

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iancullinane/sheeta/bot"
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

	err := cm.cliapp.Run(msg)
	if err != nil {
		bot.SendErrorToUser(s, err, m.ChannelID, "Deploy error")
		return
	}
	bot.SendSuccessToUser(s, m.ChannelID, "Heard cap'n")
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
	log.Println(msg)
	if msg[1] != "update" {
		return
	}

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
