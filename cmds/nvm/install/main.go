package install

import (
	"context"
	"jgttech/tools/cmds/nvm/install/lts"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "install",
		Commands: []*cli.Command{
			lts.Command(),
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
}
