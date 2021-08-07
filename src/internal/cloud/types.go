package cloud

import (
	"github.com/iancullinane/sheeta/src/internal/bot"
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

// Services are API's needed to execute module tasks
type Services struct {
	S3 S3Client
	CF CFClient
}

// StackConfig is used to generate the request
type StackConfig struct {
	Name        string                 `yaml:"name"`
	CloudConfig map[string]interface{} `yaml:"cloud-config"`
	Tags        map[string]string      `yaml:"tags"`
}
