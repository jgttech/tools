package docker

import "github.com/urfave/cli/v3"

func prune() *cli.Command {
	return &cli.Command{
		Name: "prune",
	}
}
