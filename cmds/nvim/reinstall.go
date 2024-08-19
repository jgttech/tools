package nvim

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v3"
	"jgttech/tools/sys"
)

func reinstall() *cli.Command {
	return &cli.Command{
		Name:  "reinstall",
		Usage: "Removes and re-installs the Neovim IDE",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			home, _ := os.LookupEnv("HOME")
			dirs := []string{
				path.Join(home, ".config/nvim"),
				path.Join(home, ".local/share/nvim"),
				path.Join(home, ".cache/nvim"),
			}
			cmds := []string{
				fmt.Sprintf("gh repo clone jgttech/nvim %s/.config/nvim", home),
				"npm i -g typescript typescript-language-server",
				"nvim",
			}

			for _, dir := range dirs {
				os.RemoveAll(dir)
			}

			for _, cmd := range cmds {
				sys.Catch(sys.StdCmd(cmd).Run())
			}

			return nil
		},
	}
}
