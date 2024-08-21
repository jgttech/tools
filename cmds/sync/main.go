package sync

import (
	"context"

	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "sync",
		Usage: "Performs any synchronization needed for the tools CLI.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Pulls the latest changes for the tools.
			sys.StdCatchRun("tools pull")

			// Pulls the latest changes for Neovim.
			sys.StdCatchRun("tools nvim pull")
			return nil
		},
	}
}
