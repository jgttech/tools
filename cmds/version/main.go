package version

import (
	"context"
	"fmt"

	"jgttech/tools/pkg"
	// "jgttech/tools/sys"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Displays the version of the CLI tools.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			conf := pkg.Load()
			version := conf.Version().Value()
			name := conf.Name().Value()
			// async, err := sys.AsyncRun([]string{
			// 	"npm --version",
			// 	"yarn --version",
			// 	"pnpm --version",
			// })
			//
			// if err != nil {
			// 	fmt.Printf("Error while checking tools version:\n\n")
			// 	return err
			// }
			//
			// async.Each(func(cmd *sys.AsyncCmdReturn) {
			// 	fmt.Println(cmd.StdIn)
			// })

			fmt.Printf("%s %s\n", name, version)
			return nil
		},
	}
}
