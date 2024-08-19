package docker

import (
	"context"
	"fmt"
	"jgttech/tools/sys"

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
					sys.Catch(sys.StdCmd(fmt.Sprintf("docker rm -f %s", parse("docker ps -qa"))).Run())
					return nil
				},
			},
			{
				Name:  "images",
				Usage: "Remove ALL Docker images",
				Action: func(ctx context.Context, c *cli.Command) error {
					sys.Catch(sys.StdCmd(fmt.Sprintf("docker rmi -f %s", parse("docker images -qq"))).Run())
					return nil
				},
			},
			{
				Name:  "volumes",
				Usage: "Remove ALL Docker volumes",
				Action: func(ctx context.Context, c *cli.Command) error {
					sys.Catch(sys.StdCmd(fmt.Sprintf("docker volume rm -f %s", parse("docker volume ls -q"))).Run())
					return nil
				},
			},
			{
				Name:  "networks",
				Usage: "Remove ALL Docker networks",
				Action: func(ctx context.Context, c *cli.Command) error {
					sys.Catch(sys.StdCmd(fmt.Sprintf("docker network rm -f %s", parse("docker network ls -q"))).Run())
					return nil
				},
			},
		},
	}
}
