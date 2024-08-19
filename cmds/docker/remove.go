package docker

import (
	"context"
	"fmt"
	"jgttech/tools/sys"
	"strings"

	"github.com/urfave/cli/v3"
)

func remove() *cli.Command {
	return &cli.Command{
		Name:  "remove",
		Usage: "Namespace context for removing things in Docker.",
		Commands: []*cli.Command{
			{
				Name:  "containers",
				Usage: "Remove ALL Docker containers",
				Action: func(ctx context.Context, c *cli.Command) error {
					bytes, err := sys.Cmd("docker ps -qa").CombinedOutput()
					sys.Catch(err)

					hashes := strings.TrimSpace(strings.Join(strings.Split(string(bytes), "\n"), " "))

					if hashes != "" {
						cmd := sys.StdCmd(fmt.Sprintf("docker rm -f %s", hashes))
						sys.Catch(cmd.Run())
					}

					return nil
				},
			},
			{
				Name:  "images",
				Usage: "Remove ALL Docker images",
				Action: func(ctx context.Context, c *cli.Command) error {
					// sys.Catch(sys.StdCmd("docker rmi -f $(docker images -qq)").Run())
					return nil
				},
			},
			{
				Name:  "volumes",
				Usage: "Remove ALL Docker volumes",
				Action: func(ctx context.Context, c *cli.Command) error {
					// sys.Catch(sys.Cmd("docker volume rm -f $(docker volume ls -q)").Run())
					return nil
				},
			},
			{
				Name:  "networks",
				Usage: "Remove ALL Docker networks",
				Action: func(ctx context.Context, c *cli.Command) error {
					// sys.Catch(sys.Cmd("docker network rm -f $(docker network ls -q)").Run())
					return nil
				},
			},
		},
	}
}
