package major

import (
	"context"

	"jgttech/tools/cmds/update/update"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "major",
		Usage: "Automatically increments the MAJOR version of the CLI.",
		Action: update.Repo(func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()

			conf.IncrementMajorVersion()
			conf.Write()

			sys.StdCatchRun("tools build")

			return nil
		}),
	}
}
