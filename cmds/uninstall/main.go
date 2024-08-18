package uninstall

import (
	"context"
	"fmt"
	"os"

	"jgttech/tools/path"
	"jgttech/tools/sys"

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
		Name: "uninstall",
		Action: func(ctx context.Context, _ *cli.Command) error {
			toolsPath, err := path.Join()
			sys.Panic(err)

			os.RemoveAll(toolsPath)
			fmt.Printf(message)

			return nil
		},
	}
}
