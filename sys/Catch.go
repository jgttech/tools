package sys

import "fmt"

func Catch(err error) error {
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}

	return err
}
