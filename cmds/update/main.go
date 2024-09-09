package update

import (
	"jgttech/tools/cmds/update/major"
	"jgttech/tools/cmds/update/minor"
	"jgttech/tools/cmds/update/patch"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Commands: []*cli.Command{
			major.Command(),
			minor.Command(),
			patch.Command(),
		},
	}
}
