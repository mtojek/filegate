package command

import (
	"github.com/mitchellh/cli"
)

// PullCommand performs file/folder download.
type PullCommand struct {
	ShutdownCh <-chan struct{}
	UI         cli.Ui
}

var _ cli.Command = &PullCommand{}

// Help method defines command instructions.
func (c *PullCommand) Help() string {
	return `
Usage: filegate pull [options] link_code
  ` + c.Synopsis() + `.
Options:
  --output                                          Output path to the downloaded resource.
  --signaling-server=signaling_server               Signaling endpoint, used by peers to exchange session description.
  --stun-servers=stun_server_1,stun_server_2,...    STUN servers for traversal of NAT gateways.
`
}

// Run method executes the command.
func (c *PullCommand) Run(args []string) int {

	return 0
}

// Synopsis method provides short definition.
func (c *PullCommand) Synopsis() string {
	return "Receives the shared resource using the link code"
}
