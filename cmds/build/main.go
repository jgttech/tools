package build

import (
	"context"
	"fmt"
	"os"

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
			pwd := path.Join()

			conf := pkg.Load()
			outDir := conf.OutDir().Value()
			localDir := conf.LocalDir().Value()
			versionsDir := conf.VersionsDir().Value()
			version := conf.Version().Value()
			name := conf.Name().Value()
			linkWorkingDir := path.Join(outDir, localDir)
			linkFile := path.Join(outDir, localDir, name)
			linkPath := path.Join(outDir, versionsDir, version, name)
			buildDir := path.Join(outDir, versionsDir, version, name)

			build := sys.StdCmd(fmt.Sprintf("go build -o %s", buildDir))
			build.Dir = pwd

			if err := build.Run(); err != nil {
				panic(err)
			}

			if _, err := os.Stat(linkFile); err == nil {
				os.Remove(linkFile)
			}

			ln := sys.Cmd(fmt.Sprintf("ln -s %s %s", linkPath, name))
			ln.Dir = linkWorkingDir

			// I am not worried about this throwing because
			// if the link already exists, then that is fine.
			ln.Run()

			return nil
		},
	}
}
