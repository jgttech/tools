package sys

import (
	"os/exec"
	"strings"
)

func Cmd(cmd string) *exec.Cmd {
	idx := strings.Index(cmd, " ")

	if idx == -1 {
		return nil
	}

	bin := cmd[:idx]
	argv := cmd[idx+1:]

	return exec.Command(bin, argv)
}
