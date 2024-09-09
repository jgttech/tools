package minor

import (
	"context"

	"jgttech/tools/cmds/update/update"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "minor",
		Usage: "Automatically increments the MINOR version of the CLI.",
		Action: update.Repo(func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()

			conf.IncrementMinorVersion()
			conf.Write()

			sys.StdCatchRun("tools build")

			return nil
		}),
	}
}
