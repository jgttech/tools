package sys

import (
	"os/exec"
	"strings"
)

func Cmd(cmd string) *exec.Cmd {
	args := strings.Split(cmd, " ")
	binary := args[0]
	argv := args[1:]
	command := exec.Command(binary, argv...)

	return command
}
