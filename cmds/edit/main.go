package edit

import (
	"context"
	"jgttech/tools/path"
	"jgttech/tools/sys"
	"log"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "edit",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			baseDir, err := path.Join()
			sys.Panic(err)

			command := sys.StdCmd("nvim .")
			command.Dir = baseDir

			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}
