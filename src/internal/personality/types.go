package cloud

import "github.com/urfave/cli"

// import (
// 	"github.com/iancullinane/sheeta/internal/bot"
// 	"github.com/urfave/cli/v2"
// )

// // Config keys for this package
// const (
// 	bucketNameKey = "bucketName"
// 	cloudRoleKey  = "cloudRole"
// 	regionKey     = "region"
// )

type prisoner struct {
	cliapp *cli.App
}

// // NewCloud returns a new cloud client which implements the Module interface
// func NewPrisoner() *prisoner {

// 	p := prisoner{}

// 	p.cliapp = bot.GenerateCLI(p.ExportCommands())
// 	return &p
// }
