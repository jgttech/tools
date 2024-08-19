package nvim

import (
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "nvim",
		Usage: "Neovim utilities",
		Commands: []*cli.Command{
			reinstall(),
		},
	}
}
