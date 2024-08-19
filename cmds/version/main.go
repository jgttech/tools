package version

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"jgttech/tools/.bin/env"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Displays the version of the CLI tools.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("%s %s\n", env.NAME, env.VERSION)
			return nil
		},
	}
}
