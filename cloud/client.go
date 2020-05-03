package cloud

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type cloud struct {
	r Resources
}

func NewCloud(r Resources) *cloud {
	return &cloud{
		r: r,
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (c *cloud) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if contains(m.Mentions, "sheeta") {
		tokens := strings.Split(m.ContentWithMentionsReplaced(), " ")
		if tokens[1] == "create" {
			reply := fmt.Sprintf("%s", tokens[1:])
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Creating stack: %s", reply))
			return
		}
		// reply := fmt.Sprintf("Heard! (%s)", m.Message.Content)
		s.ChannelMessageSend(m.ChannelID, "Not a valid command")
		return
	}

	// Create(c.r)

}

func contains(s []*discordgo.User, e string) bool {

	for _, a := range s {
		if a.Username == e {
			return true
		}
	}
	return false
}
