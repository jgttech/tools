package version

import (
	"context"
	"fmt"

	"jgttech/tools/pkg"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Displays the version of the CLI tools.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()
			version := conf.Version().Value()
			name := conf.Name().Value()

			fmt.Printf("%s %s\n", name, version)
			return nil
		},
	}
}
