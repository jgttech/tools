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
			commands := []string{
				"npm --version",
				"yarn --version",
				"pnpm --version",

				"nvm install --lts",
				"nvm use --lts --default",
				"npm i -g npm",
				"npm i -g yarn",
				"npm i -g pnpm",

				"npm --version",
				"yarn --version",
				"pnpm --version",
			}

			for _, cmd := range commands {
				sys.StdRun(cmd)
			}

			return nil
		},
	}
}
