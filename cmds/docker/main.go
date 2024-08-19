package docker

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "docker",
		Usage: "Docker convenience command/utilities.",
		Commands: []*cli.Command{
			ls(),
			remove(),
			kill(),
		},
	}
}
