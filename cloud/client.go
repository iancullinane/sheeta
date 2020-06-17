package cloud

import (
	"github.com/iancullinane/sheeta/bot"
	"github.com/urfave/cli/v2"
)

// Config keys for this package
const (
	bucketNameKey = "bucketName"
	cloudRoleKey  = "cloudRole"
	regionKey     = "region"
)

type cloud struct {
	s      Services
	cfg    map[string]string
	cliapp *cli.App
}

// NewCloud returns a new cloud client which implements the Module interface
func NewCloud(s Services, cfg map[string]string) *cloud {

	c := cloud{
		s:   s,
		cfg: cfg,
	}

	c.cliapp = bot.GenerateCLI(c.ExportCommands())
	return &c
}
