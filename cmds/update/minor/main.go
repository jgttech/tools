package minor

import (
	"context"

	"github.com/urfave/cli/v3"
	"jgttech/tools/cmds/update/update"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "minor",
		Usage: "Automatically increments the MINOR version of the CLI.",
		Action: update.Repo(func(ctx context.Context, cmd *cli.Command) error {
			return nil
		}),
	}
}
