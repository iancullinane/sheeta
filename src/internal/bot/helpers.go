package bot

import "github.com/bwmarrin/discordgo"

func containsUser(s []*discordgo.User, e string) bool {
	for _, a := range s {
		if a.Username == e {
			return true
		}
	}
	return false
}

func ValidateMsg(authorID string, userID string, mentions []*discordgo.User) bool {

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
