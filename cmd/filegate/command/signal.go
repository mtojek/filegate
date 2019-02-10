package command

import (
	"github.com/mitchellh/cli"
)

// SignalCommand performs file/folder upload.
type SignalCommand struct {
	ShutdownCh <-chan struct{}
	UI         cli.Ui
}

var _ cli.Command = &SignalCommand{}

// Help method defines command instructions. TODO
func (c *SignalCommand) Help() string {
	return `
Usage: filegate signal [options]
  ` + c.Synopsis() + `.
Options:
  TODO
`
}

// Run method executes the command.
func (c *SignalCommand) Run(args []string) int {

	return 0
}

// Synopsis method provides short definition.
func (c *SignalCommand) Synopsis() string {
	return "Runs the signaling server"
}
