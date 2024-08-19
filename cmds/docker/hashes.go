package docker

import (
	"jgttech/tools/sys"
	"strings"
)

func hashes(cmd string) string {
	bytes, err := sys.Cmd(cmd).CombinedOutput()
	sys.Catch(err)

	hashes := strings.TrimSpace(strings.Join(strings.Split(string(bytes), "\n"), " "))

	return hashes
}
