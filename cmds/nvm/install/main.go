package install

import (
	"github.com/urfave/cli/v3"
	"jgttech/tools/cmds/nvm/install/lts"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "install",
		Commands: []*cli.Command{
			lts.Command(),
		},
	}
}
