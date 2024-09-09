package update

import (
	"context"

	"fmt"
	"github.com/urfave/cli/v3"
	"time"

	"jgttech/tools/path"
	"jgttech/tools/sys"
)

func Repo(cb cli.ActionFunc) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		err := cb(ctx, cmd)

		if err != nil {
			return err
		}

		pwd := path.Join()

		add := sys.StdCmd(`git add .`)
		add.Dir = pwd

		commit := sys.StdCmd(fmt.Sprintf(`git commit -m "%s"`, time.Now()))
		commit.Dir = pwd

		push := sys.StdCmd(`git push`)
		push.Dir = pwd

		err = sys.Catch(add.Run())

		if err == nil {
			err = sys.Catch(commit.Run())

			if err == nil {
				sys.Catch(push.Run())
			}
		}

		return nil
	}
}
