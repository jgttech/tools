package git

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func pull() *cli.Command {
	return &cli.Command{
		Name:            "pull",
		Usage:           "A passthrough for 'git pull'.",
		SkipFlagParsing: true,
		Action: func(ctx context.Context, _ *cli.Command) error {
			argv := strings.Join(os.Args[3:], " ")
			pull := strings.TrimSpace(fmt.Sprintf(`git pull %s`, argv))
			command := sys.StdCmd(pull)
			err := command.Run()

			return err
		},
	}
}
