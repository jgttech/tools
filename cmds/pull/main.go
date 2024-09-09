package pull

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"jgttech/tools/.bin/env"
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

			sys.StdCatchRun("tools build")
			return nil
		},
	}
}
