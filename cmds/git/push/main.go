package push

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:            "push",
		Usage:           "A passthrough for 'git push'.",
		SkipFlagParsing: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			argv := strings.Join(os.Args[3:], " ")
			push := strings.TrimSpace(fmt.Sprintf(`git push %s`, argv))
			command := sys.StdCmd(push)
			err := command.Run()

			return err
		},
	}
}
