package env

import (
	"context"
	"fmt"

	"jgttech/tools/.bin/env"
	"jgttech/tools/pkg"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "env",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			envName := cmd.Args().First()
			conf := pkg.Load()
			variables := [][]string{
				{"BASE_DIR", env.BASE_DIR},
				{"NAME", conf.Name().Value()},
				{"VERSION", conf.Version().Value()},
				{"PROFILE", conf.Profile().Value()},
				{"OUT_DIR", conf.OutDir().Value()},
				{"LOCAL_DIR", conf.LocalDir().Value()},
				{"SHELL_DIR", conf.ShellDir().Value()},
				{"VERSIONS_DIR", conf.VersionsDir().Value()},
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
