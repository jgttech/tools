package nvim

import (
	"context"
	"jgttech/tools/sys"
	"os"
	"path"

	"github.com/urfave/cli/v3"
)

func sync() *cli.Command {
	return &cli.Command{
		Name:  "sync",
		Usage: "Synchronizes local and remote Neovim configuration.",
		Action: func(ctx context.Context, _ *cli.Command) error {
			home, _ := os.LookupEnv("HOME")
			dir := path.Join(home, ".config/nvim")

			cmd := sys.Cmd("git pull")
			cmd.Dir = dir

			sys.Catch(cmd.Run())
			return nil
		},
	}
}
