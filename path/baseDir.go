package path

import (
	"errors"
	"jgttech/tools/.bin/env"
	"os"
	"path"
)

func baseDir() (string, error) {
	home, ok := os.LookupEnv("HOME")

	if !ok {
		return "", errors.New("Unable to find 'HOME' environment variable.")
	}

	return path.Join(home, env.BASE_DIR), nil
}
