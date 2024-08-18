package path

import (
	p "path"
)

func Join(paths ...string) (string, error) {
	baseDir, err := baseDir()

	if err != nil {
		return "", err
	}

	return p.Join(append([]string{baseDir}, paths...)...), nil
}
