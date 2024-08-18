package edit

import (
	"context"
	"fmt"
	"jgttech/tools/path"
	"jgttech/tools/sys"
	"log"

	"github.com/urfave/cli/v3"
)

func editPath(cmd string) string {
	var arg string

	if cmd == "" {
		arg = "."
	} else {
		arg = fmt.Sprintf("cmds/%s/main.go", cmd)
	}

	return fmt.Sprintf("nvim %s", arg)
}

func Command() *cli.Command {
	return &cli.Command{
		Name: "edit",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			editCommand := cmd.Args().First()

			baseDir, err := path.Join()
			sys.Panic(err)

			command := sys.StdCmd(editPath(editCommand))
			command.Dir = baseDir

			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}
