package sys

import (
	"fmt"
	"os/exec"
	"sync"
)

type asyncCmd struct {
	in  string
	out string
	err string
	cmd *exec.Cmd
}

type asyncRuntime struct {
	cmds []*asyncCmd
}

type AsyncCmdReturn struct {
	StdIn  string
	StdOut string
	StdErr string
}

type asyncRuntimeReturn struct {
	cmds     []*AsyncCmdReturn
	AsyncErr string
	AsyncIn  []string
}

func (runtime *asyncRuntime) Start() {
	var wg sync.WaitGroup

	for _, async := range runtime.cmds {
		wg.Add(1)

		go func() {
			bytes, err := async.cmd.CombinedOutput()

			if err != nil {
				async.err = err.Error()
			} else {
				async.out = string(bytes)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

func (response *asyncRuntimeReturn) Each(callback func(*AsyncCmdReturn)) {
	for _, cmd := range response.cmds {
		callback(cmd)
	}
}

func (runtime *asyncRuntime) Return() (*asyncRuntimeReturn, bool) {
	result := asyncRuntimeReturn{}
	ok := true

	for _, async := range runtime.cmds {
		result.AsyncIn = append(result.AsyncIn, async.in)
		result.cmds = append(result.cmds, &AsyncCmdReturn{async.in, async.out, async.err})

		if async.err != "" {
			ok = false
			var msg string

			if async.err != "" {
				msg = async.err
			} else {
				msg = "--- NO ERROR(S) ---"
			}

			result.AsyncErr += fmt.Sprintf(">>> %s\n%s\n\n", async.in, msg)
		}
	}

	result.AsyncErr = "\n" + result.AsyncErr
	return &result, ok
}

func newAsyncRuntime(cmds []string) *asyncRuntime {
	runtime := asyncRuntime{}

	for _, cmd := range cmds {
		runtime.cmds = append(runtime.cmds, &asyncCmd{
			in:  cmd,
			cmd: exec.Command(cmd),
		})
	}

	return &runtime
}

func AsyncRun(cmds []string) (*asyncRuntimeReturn, error) {
	var err error

	runtime := newAsyncRuntime(cmds)
	runtime.Start()
	result, ok := runtime.Return()

	if !ok {
		err = fmt.Errorf(result.AsyncErr)
	}

	return result, err
}
