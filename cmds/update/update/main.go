package update

import (
	"context"

	"fmt"
	"time"

	"github.com/urfave/cli/v3"

	"jgttech/tools/path"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"
)

func Repo(cb cli.ActionFunc) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		conf := pkg.Load()
		oldVersion := conf.Version().Value()
		err := cb(ctx, cmd)

		if err != nil {
			return err
		}

		pwd := path.Join()
		conf = pkg.Load()
		newVersion := conf.Version().Value()

		add := sys.StdCmd(`git add .`)
		add.Dir = pwd

		commit := sys.StdCmd(fmt.Sprintf(`git commit -m "%s - %s (%s)"`, oldVersion, newVersion, time.Now()))
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
