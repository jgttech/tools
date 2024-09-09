package pull

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"jgttech/tools/.bin/env"
	"jgttech/tools/pkg"
	"jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

const (
	git_LATEST  = "Already up to date."
	git_CHANGES = "You have unstaged changes"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "pull",
		Usage: "Pull latest tools changes and compile with new changes if necessary",
		Action: func(ctx context.Context, _ *cli.Command) error {
			home, ok := os.LookupEnv("HOME")
			conf := pkg.Load()
			outDir := conf.OutDir().Value()
			versionsDir := conf.VersionsDir().Value()
			version := conf.Version().Value()
			name := conf.Name().Value()

			if !ok {
				fmt.Println("Unable to sync tools repo.")
				return nil
			}

			repo := path.Join(home, env.BASE_DIR)

			cmd := sys.Cmd("git pull")
			cmd.Dir = repo

			bytes, _ := cmd.CombinedOutput()
			out := string(bytes)

			if strings.Contains(out, git_LATEST) {
				fmt.Println(git_LATEST)
				return nil
			}

			if strings.Contains(out, git_CHANGES) {
				fmt.Println("You have unstaged changes, unable to pull the latest.")
				fmt.Println("You can use 'tools update' to commit and push your changes.")
				return nil
			}

			buildPath := path.Join(repo, outDir, versionsDir, version, name)
			build := sys.Cmd(fmt.Sprintf("go build -o %s", buildPath))
			build.Dir = repo

			sys.Catch(build.Run())
			return nil
		},
	}
}
