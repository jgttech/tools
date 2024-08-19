package docker

import "github.com/urfave/cli/v3"

func kill() *cli.Command {
	return &cli.Command{
		Name: "kill",
	}
}
