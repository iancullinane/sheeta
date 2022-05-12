package bot

import (
	"github.com/bwmarrin/discordgo"
)

func Respond(text string, i *discordgo.Interaction, d *discordgo.Session) {
	d.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			// Note: this isn't documented, but you can use that if you want to.
			// This flag just allows you to create messages visible only for the caller of the command
			// (user who triggered the command)
			Flags:   1 << 6, // ephemeral! https://discord.com/developers/docs/resources/channel#message-object-message-flags
			Content: text,
		},
	})
}

// SendSuccessToUser sends a simple 'Heard'
func SendSuccessToUser(s *discordgo.Session, channelID string, content string) {
	var me discordgo.MessageEmbed
	me.Color = 119911
	me.Fields = []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:   ":)",
			Value:  content,
			Inline: false,
		},
	}
	msgSend := discordgo.MessageSend{
		Embed: &me,
	}
	s.ChannelMessageSendComplex(channelID, &msgSend)
}

// SendErrorToUser sends the contents of an error back to the user as an
// embedded message
// TODO::Send to the user privately (does discord have ephemeral messages?)
func SendErrorToUser(s *discordgo.Session, err error, channelID string, content string) {
	errEmbed := EmbedErrorMsg(err.Error())
	msgSend := discordgo.MessageSend{
		Content: content,
		Embed:   &errEmbed,
	}
	s.ChannelMessageSendComplex(channelID, &msgSend)
}

// EmbedErrorMsg sends an embedded message with a red styled edge
func EmbedErrorMsg(s string) discordgo.MessageEmbed {
	var me discordgo.MessageEmbed
	me.Color = 14813706
	me.Description = s
	return me
}

// https://pkg.go.dev/github.com/bwmarrin/discordgo#Session.FollowupMessageCreate
// FollowupMessageCreate(interaction *Interaction, wait bool, data *WebhookParams) (*Message, error)
