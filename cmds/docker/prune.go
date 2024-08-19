package docker

import (
	"context"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func prune() *cli.Command {
	return &cli.Command{
		Name:  "prune",
		Usage: "Docker drune.",
		Action: func(ctx context.Context, _ *cli.Command) error {
			sys.Catch(sys.Cmd("docker system prune -af").Run())
			return nil
		},
	}
}
