package pkg

import (
	"fmt"
	"strings"
)

type pkgKey struct {
	name  string
	value string
	line  string
}

func (k *pkgKey) Load() {
	line := fmt.Sprintf("%s=\"%s\"", k.name, k.value)
	k.line = line
}

func (k *pkgKey) HasKey(str string) bool {
	return strings.Contains(str, k.name+"=")
}

func (k *pkgKey) Set(value string) {
	k.value = value
	k.Load()
}

func (k *pkgKey) Value() string {
	return k.value
}
