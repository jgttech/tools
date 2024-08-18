package edit

import (
	"context"
	// "jgttech/tools/sys"
	// "log"
	// "os"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name: "edit",
		Action: func(ctx context.Context, cmd *cli.Command) error {

			// cmd := sys.Cmd("nvim")
			//
			// // Bind to the current shell instance.
			// cmd.Stdout = os.Stdout
			// cmd.Stdin = os.Stdin
			// cmd.Stderr = os.Stderr
			//
			// if err := cmd.Run(); err != nil {
			// 	log.Fatal(err)
			// }

			return nil
		},
	}
}
