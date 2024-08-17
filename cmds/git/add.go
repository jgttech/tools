package git

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func add() *cli.Command {
	return &cli.Command{
		Name: "add",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("GIT ADD")
			return nil
		},
	}
}
