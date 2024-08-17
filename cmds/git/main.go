package git

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name: "git",
		Commands: []*cli.Command{
			add(),
			commit(),
			push(),
		},
	}
}
