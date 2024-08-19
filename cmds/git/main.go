package git

import "github.com/urfave/cli/v3"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "Git utilities and functions for ease of use.",
		Commands: []*cli.Command{
			add(),
			commit(),
			push(),
			pull(),
			co(),
			wip(),
			rebase(),
		},
	}
}
