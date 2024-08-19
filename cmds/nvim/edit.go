package nvim

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"os"
	"path"

	"github.com/urfave/cli/v3"
)

func edit() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Opens the Neovim config in Neovim.",
		Action: func(ctx context.Context, _ *cli.Command) error {
			home, _ := os.LookupEnv("HOME")
			editPath := path.Join(home, ".config/nvim")

			sys.Catch(sys.StdCmd(fmt.Sprintf("nvim %s", editPath)).Run())
			return nil
		},
	}
}
