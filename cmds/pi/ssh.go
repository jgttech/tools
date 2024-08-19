package pi

import (
	"context"

	"github.com/urfave/cli/v3"
	"jgttech/tools/sys"
)

func ssh() *cli.Command {
	return &cli.Command{
		Name:  "ssh",
		Usage: "Uses SSH to log into the local network pi-hole.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			sys.Catch(sys.StdCmd("ssh arez@192.168.1.112").Run())
			return nil
		},
	}
}
