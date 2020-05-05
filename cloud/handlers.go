package cloud

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/bwmarrin/discordgo"
)

const (
	bucketNameKey = "bucketName"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func (c *cloud) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !containsUser(m.Mentions, "sheeta") {
		return
	}

	ctx := context.Background()
	// Parse from the cli
	err := ParseMessage(ctx, m.Content, strings.Split(m.ContentWithMentionsReplaced(), " "))
	if err != nil {
		errEmbed := PrintEmbeddedMessage(err.Error())
		msgSend := discordgo.MessageSend{
			Embed: &errEmbed,
		}
		s.ChannelMessageSendComplex(m.ChannelID, &msgSend)
		return
	}

	input := s3.GetObjectInput{
		Bucket: aws.String(c.cfg[bucketNameKey]),
	}

	c.r.S3.GetObject(&input)
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

func PrintEmbeddedMessage(s string) discordgo.MessageEmbed {

	var me discordgo.MessageEmbed

	me.Color = 14813706
	me.Description = s

	return me
}
