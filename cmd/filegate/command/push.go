package command

import (
	"github.com/mitchellh/cli"
)

// PushCommand performs file/folder upload.
type PushCommand struct {
	ShutdownCh <-chan struct{}
	UI         cli.Ui
}

var _ cli.Command = &PushCommand{}

// Help method defines command instructions.
func (c *PushCommand) Help() string {
	return `
Usage: filegate push [options] file
  ` + c.Synopsis() + `.
Options:
  --signaling-server=signaling_server               Signaling endpoint, used by peers to exchange session description.
  --stun-servers=stun_server_1,stun_server_2,...    STUN servers for traversal of NAT gateways.
`
}

// Run method executes the command.
func (c *PushCommand) Run(args []string) int {

	return 0
}

// Synopsis method provides short definition.
func (c *PushCommand) Synopsis() string {
	return "Sends resource directly to the peer"
}
