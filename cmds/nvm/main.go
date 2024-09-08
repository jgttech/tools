package nvm

import (
	"jgttech/tools/cmds/nvm/install"
	"jgttech/tools/cmds/nvm/update"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "nvm",
		Commands: []*cli.Command{
			install.Command(),
			update.Command(),
		},
	}
}
