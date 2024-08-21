package foo

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "foo",
		Usage: "Testing foo",
		Action: func(ctx context.Context, _ *cli.Command) error {
			fmt.Println("FOO")
			return nil
		},
	}
}
