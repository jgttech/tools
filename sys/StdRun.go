package sys

func StdRun(cmd string) error {
	return StdCmd(cmd).Run()
}
