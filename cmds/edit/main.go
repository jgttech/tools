package edit

import (
	"context"
	"fmt"
	"jgttech/tools/path"
	"jgttech/tools/sys"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

const COMMANDS_DIR = "cmds"

func editPath(cmd string) string {
	var arg string
	isPath := strings.Contains(cmd, "/")
	hasGoExt := strings.Contains(cmd, ".go")

	if isPath {
		filePath, err := path.Join(cmd)
		sys.Catch(err)

		// If the path does NOT exist and it is NOT pointing
		// at a Go file, then we modify the path to include
		// the ".go" file extension.
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			if !hasGoExt {
				filePath += ".go"

				// If the path does not match anything, still, then
				// check the commands and try to see if it matches
				// anything in the commands
				if _, err = os.Stat(filePath); os.IsNotExist(err) {
					var hasCmd bool

					cmdName := strings.Split(cmd, "/")[0]
					cmdsPath, err := path.Join(COMMANDS_DIR)
					sys.Panic(err)

					dirs, err := os.ReadDir(cmdsPath)
					sys.Panic(err)

					for _, dir := range dirs {
						if cmdName == dir.Name() {
							hasCmd = true
						}
					}

					if !hasCmd {
						return fmt.Sprintf("echo \"No file found: '%s'\"", filePath)
					}

					cmdFile, err := path.Join(COMMANDS_DIR, cmd+".go")
					sys.Panic(err)

					// If the file still does not exist, then it just
					// does not exist.
					if _, err = os.Stat(cmdFile); os.IsNotExist(err) {
						return fmt.Sprintf("echo \"No file found: '%s'\"", filePath)
					}

					editCommand := fmt.Sprintf(strings.Join([]string{COMMANDS_DIR, cmd + ".go"}, "/"))

					return fmt.Sprintf("nvim %s", editCommand)
				} else {
					return fmt.Sprintf("nvim %s", cmd+".go")
				}
			}
		} else {
			return fmt.Sprintf("nvim %s", cmd)
		}
	}

	if cmd == "" {
		arg = "."
	} else {
		arg = fmt.Sprintf("cmds/%s/main.go", cmd)
	}

	return fmt.Sprintf("nvim %s", arg)
}

func Command() *cli.Command {
	return &cli.Command{
		Name:  "edit",
		Usage: "Opens the tools CLI code in Neovim. Also accepts the name of a command to open its 'main.go' file, if provided.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			editCommand := cmd.Args().First()

			baseDir, err := path.Join()
			sys.Panic(err)

			command := sys.StdCmd(editPath(editCommand))
			command.Dir = baseDir

			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}
