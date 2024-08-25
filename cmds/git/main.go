package git

import (
	"github.com/urfave/cli/v3"
	"jgttech/tools/cmds/git/add"
	"jgttech/tools/cmds/git/co"
	"jgttech/tools/cmds/git/commit"
	"jgttech/tools/cmds/git/merge"
	"jgttech/tools/cmds/git/pull"
	"jgttech/tools/cmds/git/push"
	"jgttech/tools/cmds/git/rebase"
	"jgttech/tools/cmds/git/wip"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "Git utilities and functions for ease of use.",
		Commands: []*cli.Command{
			add.Command(),
			co.Command(),
			commit.Command(),
			merge.Command(),
			pull.Command(),
			push.Command(),
			rebase.Command(),
			wip.Command(),
		},
	}
}
