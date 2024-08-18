package main

import (
	"context"
	"jgttech/tools/cmds/build"
	"jgttech/tools/cmds/edit"
	"jgttech/tools/cmds/git"
	"jgttech/tools/cmds/sync"
	"jgttech/tools/cmds/test"
	"jgttech/tools/cmds/update"
	"jgttech/tools/cmds/version"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	tools := &cli.Command{
		Name: "tools",
		Commands: []*cli.Command{
			build.Command(),
			edit.Command(),
			git.Command(),
			sync.Command(),
			test.Command(),
			update.Command(),
			version.Command(),
		},
	}

	if err := tools.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
