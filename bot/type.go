package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli/v2"
)

// Handler is the function definition and its flags
type Handler struct {
	Fn    func(s *discordgo.Session, m *discordgo.MessageCreate)
	Flags []cli.Flag
}
