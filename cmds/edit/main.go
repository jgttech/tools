package edit

import (
	"context"
	"errors"
	"jgttech/tools/.bin/env"
	"log"
	"os"
	"os/exec"
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

			command := exec.Command("nvim")
			command.Dir = path.Join(home, env.BASE_DIR)

			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}
