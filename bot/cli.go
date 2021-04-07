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
func GenerateCLI(actions []Action) *cli.App {

	newCLI := buildCLI(actions)
	newCLI.HideHelp = true
	newCLI.HideHelpCommand = true

	return newCLI
}

func buildCLI(actions []Action) *cli.App {

	var tmpCLI cli.App
	var cmds []*cli.Command
	for _, hand := range actions {
		cmds = append(cmds, buildCmd(hand, ""))
	}
	tmpCLI.Commands = cmds
	return &tmpCLI
}

func buildCmd(a Action, usage string) *cli.Command {

	c := cli.Command{
		Name:            a.Name,
		Usage:           usage,
		Flags:           a.Flags,
		Action:          a.APIFn,
		HideHelp:        true,
		HideHelpCommand: true,
	}

	return &c
}
