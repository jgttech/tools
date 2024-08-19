package docker

import (
	"context"
	"fmt"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func kill() *cli.Command {
	return &cli.Command{
		Name:  "kill",
		Usage: "Kills ALL Docker processes.",
		Action: func(ctx context.Context, c *cli.Command) error {
			sys.Catch(sys.Cmd(fmt.Sprintf("docker kill %s", hashes("docker ps -qa"))).Run())
			return nil
		},
	}
}
