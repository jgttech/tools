package git

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
	"jgttech/tools/sys"
)

func wip() *cli.Command {
	var signed bool

	return &cli.Command{
		Name:  "wip",
		Usage: "Automatically generates a commit formatted as 'WIP <timestamp>'",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "signed",
				Aliases:     []string{"S"},
				Destination: &signed,
			},
		},
		Action: func(ctx context.Context, _ *cli.Command) error {
			var signedFlag string
			var commitMessage []string

			if signed {
				signedFlag = "-S"
			}

			timestamp := time.Now()
			commitTokens := []string{"git commit", signedFlag, fmt.Sprintf(`-m "WIP %s"`, timestamp)}

			for _, token := range commitTokens {
				if strings.TrimSpace(token) != "" {
					commitMessage = append(commitMessage, token)
				}
			}

			add := sys.StdCmd(`git add .`)
			commit := sys.StdCmd(strings.Join(commitMessage, " "))
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
