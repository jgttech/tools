package main

import (
	"context"
	"jgttech/tools/cmds/install"
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
			install.Command(),
		},
	}

	if err := tools.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
