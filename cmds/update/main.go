package update

import (
	"context"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			pwd, err := path.Join()
			sys.Panic(err)

			add := sys.StdCmd("git add .")
			add.Dir = pwd

			// commit := sys.StdCmd(fmt.Sprintf("git commit -m \"WIP %s\"", time.Now()))
			// commit.Dir = pwd

			// push := sys.StdCmd("git push")
			// push.Dir = pwd

			if err := add.Run(); err == nil {
				// if err = commit.Run(); err == nil {
				// 	if err = push.Run(); err == nil {
				// 		fmt.Println("Done")
				// 	}
				// }
			}

			return err
		},
	}
}
