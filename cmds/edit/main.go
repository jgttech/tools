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
			var editPath []string

			for idx := range cmd.Args().Len() {
				arg := cmd.Args().Get(idx)
				tokens := strings.Split(arg, " ")

				for idx, token := range tokens {
					tokens[idx] = strings.TrimSpace(token)
				}

				editPath = append(editPath, strings.Join(tokens, "/"))
			}

			var target string
			editCommand := strings.TrimSpace(strings.Join(editPath, "/"))

			rawEditPath := path.Join(editCommand)
			cmdEditPath := path.Join("cmds", editCommand)
			rawEditPathMain := strings.Join([]string{rawEditPath, "main.go"}, "/")
			cmdEditPathMain := strings.Join([]string{cmdEditPath, "main.go"}, "/")

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

			if target == "" {
				fmt.Println(fmt.Sprintf("No file found for comamnd: '%s'", editCommand))
				return nil
			}

			sys.StdCatchRun(fmt.Sprintf("nvim %s", target))
			return nil
		},
	}
}
