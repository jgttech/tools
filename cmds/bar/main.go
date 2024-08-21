package bar

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "bar",
		Usage: "Testing bar",
		Action: func(ctx context.Context, _ *cli.Command) error {
			fmt.Println("BAR`")
			return nil
		},
	}
}
