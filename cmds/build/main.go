package build

import (
	"context"
	"fmt"

	"jgttech/tools/path"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "build",
		Usage: "Builds/Rebuilds the tools CLI",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()
			pwd := path.Join()
			outDir := conf.OutDir().Value()
			versionsDir := conf.VersionsDir().Value()
			version := conf.Version().Value()
			name := conf.Name().Value()
			buildDir := path.Join(outDir, versionsDir, version, name)

			build := sys.StdCmd(fmt.Sprintf("go build -o %s", buildDir))
			build.Dir = pwd

			if err := build.Run(); err != nil {
				sys.Panic(err)
			}

			return nil
		},
	}
}
