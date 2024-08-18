package git

import (
	"context"
	"jgttech/tools/sys"
	"strings"

	"github.com/urfave/cli/v3"
)

func push() *cli.Command {
	var force bool

	return &cli.Command{
		Name: "push",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Destination: &force,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			push := "git push"

			if force {
				push += " -f"
			}

			command := sys.StdCmd(strings.TrimSpace(push))
			err := command.Run()

			return err
		},
	}
}
