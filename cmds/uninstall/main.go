package uninstall

import (
	"context"
	"fmt"
	"os"

	"jgttech/tools/path"

	"github.com/urfave/cli/v3"
)

const message string = `
| WARNING
|
| The link that was setup in your
| ZSH configuration was NOT removed.
| You will have to manually remove
| that if you no longer want the link
| to exist.
`

func Command() *cli.Command {
	return &cli.Command{
		Name:  "uninstall",
		Usage: "Removes the tools CLI source code, but NOT the shell profile link.",
		Action: func(ctx context.Context, _ *cli.Command) error {
			toolsPath := path.Join()

			os.RemoveAll(toolsPath)
			fmt.Printf(message)

			return nil
		},
	}
}
