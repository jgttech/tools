package edit

import (
	"context"
	"fmt"
	"os"
	"strings"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Opens the tools CLI code in Neovim. Also accepts the name of a command to open its 'main.go' file, if provided.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var target string
			editCommand := cmd.Args().First()

			rawEditPath, err := path.Join(editCommand)
			sys.Panic(err)

			cmdEditPath, err := path.Join("cmds", editCommand)
			sys.Panic(err)

			rawEditPathMain := strings.Join([]string{rawEditPath, "main.go"}, "/")
			sys.Panic(err)

			cmdEditPathMain := strings.Join([]string{cmdEditPath, "main.go"}, "/")
			sys.Panic(err)

			editPermutations := []string{
				rawEditPath,
				cmdEditPath,
				rawEditPath + ".go",
				cmdEditPath + ".go",
				rawEditPathMain,
				cmdEditPathMain,
			}

			for _, filePath := range editPermutations {
				stat, err := os.Stat(filePath)

				if err == nil && !stat.IsDir() {
					target = filePath
					break
				}
			}

			sys.StdCatchRun(fmt.Sprintf("nvim %s", target))
			return nil
		},
	}
}
