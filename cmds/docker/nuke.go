package docker

import "github.com/urfave/cli/v3"

func nuke() *cli.Command {
	return &cli.Command{
		Name: "nuke",
	}
}
