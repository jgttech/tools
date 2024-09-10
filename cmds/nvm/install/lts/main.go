package lts

import (
	"context"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "lts",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			sys.StdRun("nvm install --lts")
			sys.StdRun("nvm use --lts --default")
			sys.StdRun("npm i -g npm")
			sys.StdRun("npm i -g yarn")
			sys.StdRun("npm i -g pnpm")
			return nil
		},
	}
}
