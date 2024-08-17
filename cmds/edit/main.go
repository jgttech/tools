package edit

import (
	"context"
	// "os/exec"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "edit",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// exec.Command("nvim")
			return nil
		},
	}
}
