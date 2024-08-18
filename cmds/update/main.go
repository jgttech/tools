package update

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func runFromPath(bin string, wd string) func(cmd string) *exec.Cmd {
	return func(args string) *exec.Cmd {
		exe := fmt.Sprintf(`%s %s`, bin, args)
		// fmt.Printf(`COMMAND: [%s]\n`, exe)

		cmd := sys.StdCmd(exe)
		cmd.Dir = wd

		return cmd
	}
}

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, _ *cli.Command) error {
			timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			pwd, err := path.Join()
			sys.Panic(err)

			git := runFromPath("git", pwd)
			add := git(`add .`)
			commit := git(fmt.Sprintf(`commit -m "WIP %s"`, timestamp))
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
