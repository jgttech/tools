package pi

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "pi",
		Usage: "Pi-hole utilities.",
		Commands: []*cli.Command{
			ssh(),
		},
	}
}
