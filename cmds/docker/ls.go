package docker

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

type proc struct {
	header string
	exec   *exec.Cmd
}

func (p *proc) Run() error {
	return sys.Catch(p.exec.Run())
}

func ls() *cli.Command {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	bold := style.Bold(true).Render

	return &cli.Command{
		Name:  "ls",
		Usage: "Same as 'ls' but lists the docker state (ds).",
		Action: func(ctx context.Context, _ *cli.Command) error {
			commands := []proc{
				{
					"CONTAINERS",
					sys.StdCmd(`docker ps -a`),
				},
				{
					"IMAGES",
					sys.StdCmd(`docker images -a`),
				},
				{
					"VOLUMES",
					sys.StdCmd(`docker volume ls`),
				},
				{
					"NETWORKS",
					sys.StdCmd(`docker network ls`),
				},
			}

			for _, command := range commands {
				fmt.Println(bold("--------------------------------"))
				fmt.Println(bold(fmt.Sprintf("%s\n\n", command.header)))
				err := command.Run()
				fmt.Printf("\n\n")

				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
