package build

import (
	"context"
	"fmt"
	"os"

	"jgttech/tools/list"
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
			versionPath := path.Join(outDir, versionsDir)
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

			files, err := os.ReadDir(versionPath)

			// Reverse the list of files because the natural
			// order of the list should suffice to know which
			// versions to remove, as long as it is not the
			// current version.
			files = list.Reverse(files)

			maxSize := 4
			isMaxSize := len(files) >= maxSize

			if err != nil {
				panic(err)
			}

			if isMaxSize {
				for idx, file := range files {
					if version == file.Name() {
						continue
					}

					if idx >= maxSize {
						err := os.RemoveAll(path.Join(outDir, versionsDir, file.Name()))

						if err != nil {
							fmt.Println(err.Error())
						}
					}
				}
			}

			return nil
		},
	}
}
