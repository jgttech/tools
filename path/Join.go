package path

import (
	p "path"
)

func Join(paths ...string) string {
	baseDir, err := baseDir()

	if err != nil {
		panic(err)
	}

	return p.Join(append([]string{baseDir}, paths...)...)
}
