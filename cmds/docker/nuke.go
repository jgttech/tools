package docker

import (
	"context"
	"fmt"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func nuke() *cli.Command {
	return &cli.Command{
		Name:  "nuke",
		Usage: "Runs all the removal operations for Docker containers.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cmds := []string{"containers", "images", "volumes", "networks"}

			for _, cmd := range cmds {
				sys.Catch(sys.StdCmd(fmt.Sprintf("tools docker remove %s", cmd)).Run())
			}

			sys.Catch(sys.StdCmd("tools docker ls").Run())
			return nil
		},
	}
}
