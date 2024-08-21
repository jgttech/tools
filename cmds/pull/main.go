package pull

import (
	"context"
	"fmt"
	"os"
	"path"

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
			fmt.Println(string(bytes))

			// cmd := sys.Cmd("git pull")
			// cmd.Dir = repo
			//
			// bytes, err := cmd.Output()
			// sys.Catch(err)
			//
			// out := string(bytes)
			//
			// if strings.Contains(out, "Already up to date.") {
			// 	return nil
			// }

			// buildPath := path.Join(repo, env.OUT_DIR, env.VERSIONS_DIR, env.VERSION, env.NAME)
			// build := sys.Cmd(fmt.Sprintf("go build -o %s", buildPath))
			// build.Dir = repo

			// fmt.Println(fmt.Sprintf("go build -o %s", buildPath))
			return nil
		},
	}
}
