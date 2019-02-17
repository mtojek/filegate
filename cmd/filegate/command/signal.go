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
	// TODO Start HTTP server
	// TODO Wait for QUIT signal
	// TODO On sender connect:
	// TODO     1. Sender gives his public key: create a link code ~chat room
	// TODO		2. Destroy a link code if the sender disconnects
	// TODO     3. Receiver requests resource with a link-code, give back the sender's public key
	// TODO		4. Server proxies encrypted "chat"
	return 0
}

// Synopsis method provides short definition.
func (c *SignalCommand) Synopsis() string {
	return "Runs the signaling server"
}
