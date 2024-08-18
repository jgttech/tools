package edit

import (
	"context"
	"errors"
	"fmt"
	"jgttech/tools/.bin/env"
	"os"
	"path"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "edit",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			home, ok := os.LookupEnv("HOME")

			if !ok {
				return errors.New("Unable to find 'HOME' environment variable.")
			}

			baseDir := path.Join(home, env.BASE_DIR)
			fmt.Println("BASE_DIR:", baseDir)

			return nil
		},
	}
}
