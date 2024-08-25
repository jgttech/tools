package add

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
		Name:            "add",
		Usage:           "A passthrough for 'git add'.",
		SkipFlagParsing: true,
		Action: func(ctx context.Context, _ *cli.Command) error {
			argv := strings.Join(os.Args[3:], " ")
			add := strings.TrimSpace(fmt.Sprintf(`git add %s`, argv))
			cmd := sys.StdCmd(add)
			err := cmd.Run()

			return err
		},
	}
}
