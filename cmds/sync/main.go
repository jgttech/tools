package sync

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "sync",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
}
