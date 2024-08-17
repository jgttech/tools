package git

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func push() *cli.Command {
	return &cli.Command{
		Name:     "push",
		Category: "git",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("GIT PUSH")
			return nil
		},
	}
}
