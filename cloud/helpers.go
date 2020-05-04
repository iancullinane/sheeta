package cloud

import "github.com/bwmarrin/discordgo"

func containsUser(s []*discordgo.User, e string) bool {
	for _, a := range s {
		if a.Username == e {
			return true
		}
	}
	return false
}
