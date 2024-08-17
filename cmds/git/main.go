package git

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Commands: []*cli.Command{
			add(),
			commit(),
			push(),
		},
	}
}
