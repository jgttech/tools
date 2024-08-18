package sys

import (
	"os"
	"os/exec"
)

func StdCmd(cmd string) *exec.Cmd {
	command := Cmd(cmd)

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	return command
}
