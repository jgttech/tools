package test

import (
	"context"

	"github.com/urfave/cli/v3"
	"jgttech/tools/sys"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "test",
		Action: func(ctx context.Context, _ *cli.Command) error {
			sys.StdCmd(`git commit -m "WIP 143255735452432"`)

			return nil
		},
	}
}
