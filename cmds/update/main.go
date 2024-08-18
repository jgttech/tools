package update

import (
	"context"
	"fmt"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, _ *cli.Command) error {
			pwd, err := path.Join()
			sys.Panic(err)

			add := sys.StdCmd("git add .")
			add.Dir = pwd

			err = add.Run()
			sys.Panic(err)

			// timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			cmd := `git commit -m "WIP"`
			commit := sys.StdCmd(cmd)
			commit.Dir = pwd

			fmt.Println("COMMAND:", cmd)

			err = commit.Run()
			sys.Panic(err)

			// push := sys.StdCmd("git push")
			// push.Dir = pwd

			return err
		},
	}
}
