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

			out, err := cmd.CombinedOutput()
			sys.Catch(err)

			fmt.Printf(string(out))
			return nil
		},
	}
}
