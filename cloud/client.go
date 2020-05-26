package cloud

import (
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
	return &cloud{
		s:   s,
		cfg: cfg,
	}
}
