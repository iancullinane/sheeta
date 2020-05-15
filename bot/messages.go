package bot

import "github.com/bwmarrin/discordgo"

// SendSuccessToUser sends a simple 'Heard'
func SendSuccessToUser(s *discordgo.Session, channelID string, content string) {
	var me discordgo.MessageEmbed
	me.Color = 119911
	me.Fields = []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:  ":)",
			Value: content,
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
