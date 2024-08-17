package main

import (
	"context"
	"jgttech/tools/cmds/git"
	"jgttech/tools/cmds/sync"
	"jgttech/tools/cmds/version"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	tools := &cli.Command{
		Name: "tools",
		Commands: []*cli.Command{
			version.Command(),
			sync.Command(),
			git.Command(),
		},
	}

	if err := tools.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
