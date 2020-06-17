package bot

import (
	"encoding/json"

	"github.com/urfave/cli/v2"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// GenerateCLI returns a CLI app to be used for interpreting chatops
func GenerateCLI(handlers []Command) *cli.App {

	newCLI := buildCLI(handlers)
	newCLI.HideHelp = true
	newCLI.HideHelpCommand = true

	return newCLI
}

func buildCLI(handlers []Command) *cli.App {

	var tmpCLI cli.App
	var cmds []*cli.Command
	for _, hand := range handlers {
		cmds = append(cmds, buildCmd(hand, ""))
	}
	tmpCLI.Commands = cmds
	return &tmpCLI
}

func buildCmd(handler Command, usage string) *cli.Command {

	c := cli.Command{
		Name:            handler.Name,
		Usage:           usage,
		Flags:           handler.Flags,
		Action:          handler.APIFn,
		HideHelp:        true,
		HideHelpCommand: true,
	}

	return &c
}
