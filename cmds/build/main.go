package build

import (
	"context"
	"fmt"

	"jgttech/tools/.bin/env"
	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "build",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			pwd, err := path.Join()
			sys.Panic(err)

			buildDir, err := path.Join(env.OUT_DIR, env.VERSIONS_DIR, env.VERSION, env.NAME)
			sys.Panic(err)

			build := sys.StdCmd(fmt.Sprintf("go build -o %s", buildDir))
			build.Dir = pwd

			if err := build.Run(); err != nil {
				sys.Panic(err)
			}

			return nil
		},
	}
}
