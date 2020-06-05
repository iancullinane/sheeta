package bot

import "github.com/bwmarrin/discordgo"

// Config keys for this package
const (
	bucketNameKey = "bucketName"
	cloudRoleKey  = "cloudRole"
	regionKey     = "region"
)

type bot struct {
	modules []Module
}

// NewBot returns a new bot client
func NewBot() *bot {
	return &bot{}
}

// Module is an independent set of actions containing its cli and handlers
type Module interface {
	GenerateCLI()
	ExportHandlers() []func(s *discordgo.Session, m *discordgo.MessageCreate)
}

func (b *bot) AddModule(m Module) error {
	b.modules = append(b.modules, m)
	return nil
}

func (b *bot) GetModules() []Module {
	return b.modules
}
