package lts

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "lts",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			println("NVM INSTALL LTS")
			return nil
		},
	}
}
