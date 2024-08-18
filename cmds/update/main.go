package update

import (
	"context"
	"fmt"
	"time"

	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func catch(err error) error {
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}

	return nil
}

func Command() *cli.Command {
	return &cli.Command{
		Name: "update",
		Action: func(ctx context.Context, _ *cli.Command) error {
			pwd, err := path.Join()
			sys.Panic(err)

			add := sys.StdCmd(`git add .`)
			add.Dir = pwd

			commit := sys.StdCmd(fmt.Sprintf(`git commit -m "WIP %s"`, time.Now()))
			commit.Dir = pwd

			push := sys.StdCmd(`git push`)
			push.Dir = pwd

			err = add.Run()
			err = catch(err)

			err = commit.Run()
			err = catch(err)

			err = push.Run()
			err = catch(err)

			return err
		},
	}
}
