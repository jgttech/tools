package patch

import (
	"context"

	"jgttech/tools/cmds/update/update"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "patch",
		Usage: "Automatically increments the PATCH version of the CLI.",
		Action: update.Repo(func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()

			conf.IncrementPatchVersion()
			conf.Write()

			sys.StdCatchRun("tools build")

			return nil
		}),
	}
}
