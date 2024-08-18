package git

import (
	"context"
	"fmt"
	"strings"

	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func add() *cli.Command {
	return &cli.Command{
		Name: "add",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()
			command := sys.StdCmd(fmt.Sprintf("git add %s", strings.Join(args, " ")))
			err := command.Run()

			return err
		},
	}
}
