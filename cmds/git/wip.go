package git

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"time"

	"github.com/urfave/cli/v3"
)

func wip() *cli.Command {
	var signed bool

	return &cli.Command{
		Name:            "wip",
		SkipFlagParsing: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "signed",
				Aliases:     []string{"S"},
				Destination: &signed,
			},
		},
		Action: func(ctx context.Context, _ *cli.Command) error {
			var signedFlag string

			if signed {
				signedFlag = "-S"
			}

			timestamp := time.Now()
			add := sys.StdCmd(`git add .`)
			commit := sys.StdCmd(fmt.Sprintf(`git commit %s "WIP %s"`, signedFlag, timestamp))
			push := sys.StdCmd(`git push`)

			err := sys.Catch(add.Run())

			if err == nil {
				err = sys.Catch(commit.Run())

				if err == nil {
					sys.Catch(push.Run())
				}
			}

			return err
		},
	}
}
