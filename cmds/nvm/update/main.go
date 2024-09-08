package update

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			println("NVM UPDATE")
			return nil
		},
	}
}
