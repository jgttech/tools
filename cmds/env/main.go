package env

import (
	"context"
	"fmt"

	"jgttech/tools/.bin/env"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "env",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			envName := cmd.Args().First()
			variables := [][]string{
				{"NAME", env.NAME},
				{"VERSION", env.VERSION},
				{"PROFILE", env.PROFILE},
				{"OUT_DIR", env.OUT_DIR},
				{"BASE_DIR", env.BASE_DIR},
				{"LOCAL_DIR", env.LOCAL_DIR},
				{"SHELL_DIR", env.SHELL_DIR},
				{"VERSIONS_DIR", env.VERSIONS_DIR},
			}

			if envName != "" {
				for _, variable := range variables {
					if variable[0] == envName {
						fmt.Println(variable[1])
						return nil
					}
				}

				return nil
			}

			for _, variable := range variables {
				fmt.Println(fmt.Sprintf("%s='%s'", variable[0], variable[1]))
			}

			return nil
		},
	}
}
