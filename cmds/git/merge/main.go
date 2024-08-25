package merge

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "merge",
		Usage: "Automatically perform a Git merge process.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			branch := cmd.Args().First()

			bytes, err := sys.Cmd("git rev-parse --abbrev-ref HEAD").CombinedOutput()
			sys.Catch(err)

			current := string(bytes)

			if branch == "" {
				branch = "main"
			}

			cmds := []string{
				fmt.Sprintf("git checkout %s", branch),
				"git pull",
				fmt.Sprintf("git checkout %s", current),
				fmt.Sprintf("git merge %s", branch),
			}

			for _, cmd := range cmds {
				sys.Catch(sys.StdCmd(strings.TrimSpace(cmd)).Run())
			}

			return nil
		},
	}
}
