package git

import (
	"context"
	"fmt"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func rebase() *cli.Command {
	return &cli.Command{
		Name:  "rebase",
		Usage: "Passthrough for Git rebase, which defaults to 'main'",
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
				fmt.Sprintf("git rebase %s", branch),
			}

			for _, cmd := range cmds {
				sys.Catch(sys.StdCmd(cmd).Run())
			}

			return nil
		},
	}
}
