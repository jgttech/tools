package update

import (
	"context"
	"fmt"
	"os/exec"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func runFromPath(bin string, wd string) func(cmd string) *exec.Cmd {
	return func(args string) *exec.Cmd {
		cmd := sys.StdCmd(fmt.Sprintf(`%s %s`, bin, args))
		cmd.Dir = wd

		return cmd
	}
}

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, _ *cli.Command) error {
			pwd, err := path.Join()
			sys.Panic(err)

			// timestamp := strconv.FormatInt(time.Now().Unix(), 10)

			git := runFromPath("git", pwd)

			add := git(`add .`)
			commit := git(`commit -m "WIP"`)
			push := git(`push`)

			err = add.Run()
			sys.Panic(err)

			err = commit.Run()
			sys.Panic(err)

			err = push.Run()
			sys.Panic(err)

			return err
		},
	}
}
