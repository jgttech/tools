package build

import (
	"context"
	"fmt"

	"jgttech/tools/.bin/env"
	"jgttech/tools/path"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

// go build -o $HOME/.tools/.bin/versions/1.0.0/tools
func Command() *cli.Command {
	return &cli.Command{
		Name: "build",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			_, err := path.Join()
			sys.Panic(err)

			buildDir, err := path.Join(env.OUT_DIR, env.VERSIONS_DIR, env.VERSION, env.NAME)
			sys.Panic(err)

			fmt.Println(fmt.Sprintf("go build -o %s", buildDir))

			// build := sys.StdCmd(fmt.Sprintf("go build -o %s", buildDir))
			// build.Dir = pwd
			//
			return nil
		},
	}
}
