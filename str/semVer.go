package str

import (
	"fmt"
	"strconv"
	"strings"
)

type semVer struct {
	version string
	major   int
	minor   int
	patch   int
}

const (
	MAJOR = 0
	MINOR = 1
	PATCH = 2
)

func (s *semVer) sync(mode int) {
	switch mode {
	case MAJOR:
		s.version = fmt.Sprintf("%d.0.0", s.major)
		break
	case MINOR:
		s.version = fmt.Sprintf("%d.%d.0", s.major, s.minor)
		break
	case PATCH:
		s.version = fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
		break
	}
}

func (s *semVer) Load(version string) {
	version = strings.ReplaceAll(strings.TrimSpace(version), "\"", "")
	tokens := strings.Split(version, ".")
	major, err := strconv.Atoi(tokens[0])

	if err != nil {
		panic(err)
	}

	minor, err := strconv.Atoi(tokens[1])

	if err != nil {
		panic(err)
	}

	patch, err := strconv.Atoi(tokens[2])

	if err != nil {
		panic(err)
	}

	s.major = major
	s.minor = minor
	s.patch = patch
	s.version = fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func (s *semVer) IncrementPatchBy(by int) {
	s.patch += by
	s.sync(PATCH)
}

func (s *semVer) IncrementMinorBy(by int) {
	s.minor += by
	s.sync(MINOR)
}

func (s *semVer) IncrementMajorBy(by int) {
	s.major += by
	s.sync(MAJOR)
}

func (s *semVer) Value() string {
	return s.version
}
