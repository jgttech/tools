package main

import (
	"context"
	"jgttech/tools/cmds/build"
	"jgttech/tools/cmds/docker"
	"jgttech/tools/cmds/edit"
	"jgttech/tools/cmds/env"
	"jgttech/tools/cmds/git"
	"jgttech/tools/cmds/nvim"
	"jgttech/tools/cmds/pi"
	"jgttech/tools/cmds/sync"
	"jgttech/tools/cmds/uninstall"
	"jgttech/tools/cmds/update"
	"jgttech/tools/cmds/version"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	tools := &cli.Command{
		Name:        "tools",
		Description: "Custom tools for my own purposes.",
		Commands: []*cli.Command{
			build.Command(),
			docker.Command(),
			edit.Command(),
			env.Command(),
			git.Command(),
			nvim.Command(),
			pi.Command(),
			sync.Command(),
			uninstall.Command(),
			update.Command(),
			version.Command(),
		},
	}

	if err := tools.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
