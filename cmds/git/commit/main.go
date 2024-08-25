package commit

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
		Name:            "commit",
		Usage:           "A passthrough for 'git commit'.",
		SkipFlagParsing: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			argv := strings.Join(os.Args[3:], " ")
			commit := strings.TrimSpace(fmt.Sprintf(`git commit %s`, argv))
			command := sys.StdCmd(commit)
			err := command.Run()

			return err
		},
	}
}
