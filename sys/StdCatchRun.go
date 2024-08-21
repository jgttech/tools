package sys

func StdCatchRun(cmd string) error {
	return Catch(StdRun(cmd))
}
