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

			if strings.Contains(out, "Already up to date.") {
				fmt.Println("HELLO WORLD")
				return nil
			}

			if strings.Contains(out, "You have unstaged changes") {
				fmt.Println("You have unstaged changes, unable to pull the latest.")
				return nil
			}

			buildPath := path.Join(repo, env.OUT_DIR, env.VERSIONS_DIR, env.VERSION, env.NAME)
			// build := sys.Cmd(fmt.Sprintf("go build -o %s", buildPath))
			// build.Dir = repo

			fmt.Println(fmt.Sprintf("go build -o %s", buildPath))
			return nil
		},
	}
}
