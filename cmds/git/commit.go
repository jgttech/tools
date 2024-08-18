package git

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"strings"

	"github.com/urfave/cli/v3"
)

func commit() *cli.Command {
	var msg string
	var sign bool

	return &cli.Command{
		Name: "commit",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "message",
				Aliases:     []string{"m"},
				Destination: &msg,
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "sign",
				Aliases:     []string{"S"},
				Destination: &sign,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			commit := "git commit"

			if msg != "" {
				commit = commit + fmt.Sprintf(` -m "%s"`, msg)
			}

			if sign {
				commit = commit + " -S"
			}

			fmt.Println(strings.TrimSpace(commit))

			command := sys.StdCmd(commit)
			err := command.Run()

			return err
		},
	}
}
