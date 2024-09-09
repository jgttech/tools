package pkg

import (
	"fmt"
	"jgttech/tools/.bin/env"
	"jgttech/tools/path"
	"jgttech/tools/str"
	"os"
	"strings"
)

type pkg struct {
	path        string
	file        []string
	version     *pkgKey
	name        *pkgKey
	profile     *pkgKey
	outDir      *pkgKey
	localDir    *pkgKey
	versionsDir *pkgKey
	shellDir    *pkgKey
}

func (conf *pkg) load() {
	var file []string

	pkgPath := path.Join("pkg.sh")
	pkgFile, err := os.ReadFile(pkgPath)

	if err != nil {
		panic(err)
	}

	conf.path = pkgPath
	lines := strings.Split(string(pkgFile), "\n")

	for _, line := range lines {
		file = append(file, line)
		ln := strings.ReplaceAll(line, "\"", "")

		if strings.Contains(line, "=") {
			tokens := strings.Split(ln, "=")
			key := tokens[0]
			value := tokens[1]

			confKey := &pkgKey{}
			confKey.value = value

			if key == "VERSION" {
				confKey.name = "VERSION"
				conf.version = confKey
			} else if key == "NAME" {
				confKey.name = "NAME"
				conf.name = confKey
			} else if key == "PROFILE" {
				confKey.name = "PROFILE"
				conf.profile = confKey
			} else if key == "OUT_DIR" {
				confKey.name = "OUT_DIR"
				conf.outDir = confKey
			} else if key == "LOCAL_DIR" {
				confKey.name = "LOCAL_DIR"
				conf.localDir = confKey
			} else if key == "VERSIONS_DIR" {
				confKey.name = "VERSIONS_DIR"
				conf.versionsDir = confKey
			} else if key == "SHELL_DIR" {
				confKey.name = "SHELL_DIR"
				conf.shellDir = confKey
			}

			// Sets the internal "line" property to represent
			// that line in the file and keep an updated copy
			// so that saving the changes to the config can be
			// done from self-maintained lines within each "pkgKey".
			//
			// The idea is just to key each key update, as needed,
			// and when the changes need to be written to the file
			// system, then we just read the "line" from each
			// config file property.
			confKey.Load()
		}
	}

	conf.file = file
}

func (p *pkg) Version() *pkgKey {
	return p.version
}

func (p *pkg) Name() *pkgKey {
	return p.name
}

func (p *pkg) Profile() *pkgKey {
	return p.profile
}

func (p *pkg) OutDir() *pkgKey {
	return p.outDir
}

func (p *pkg) LocalDir() *pkgKey {
	return p.localDir
}

func (p *pkg) VersionsDir() *pkgKey {
	return p.versionsDir
}

func (p *pkg) ShellDir() *pkgKey {
	return p.shellDir
}

func (p *pkg) sync(key, update string) {
	for idx, line := range p.file {
		if strings.Contains(line, key+"=") {
			p.file[idx] = update
		}
	}
}

func (p *pkg) incrementVersion(mode, by int) {
	semver := str.NewSemVer(p.version.value)

	switch mode {
	case str.MAJOR:
		semver.IncrementMajorBy(by)
		break
	case str.MINOR:
		semver.IncrementMinorBy(by)
		break
	case str.PATCH:
		semver.IncrementPatchBy(by)
		break
	}

	p.version.Set(semver.Value())
	p.sync(p.version.name, p.version.line)
}

func (p *pkg) IncrementPatchVersion() {
	p.incrementVersion(str.PATCH, 1)
}

func (p *pkg) IncrementMinorVersion() {
	p.incrementVersion(str.MINOR, 1)
}

func (p *pkg) IncrementMajorVersion() {
	p.incrementVersion(str.MAJOR, 1)
}

func (p *pkg) Write() {
	var data string

	for _, line := range p.file {
		data += line + "\n"
	}

	data = strings.TrimSpace(data)
	os.WriteFile(p.path, []byte(data), 0666)
}

func (p *pkg) GenerateEnv() {
	envFilePath := path.Join(p.outDir.value, "env/env.go")
	data := "package env\n\n"
	data += "const (\n"

	data += fmt.Sprintf("  %s = \"%s\"\n", "BASE_DIR", env.BASE_DIR)
	data += fmt.Sprintf("  %s = \"%s\"\n", "OUT_DIR", p.outDir.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "VERSIONS_DIR", p.versionsDir.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "LOCAL_DIR", p.localDir.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "SHELL_DIR", p.shellDir.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "VERSION", p.version.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "NAME", p.name.value)
	data += fmt.Sprintf("  %s = \"%s\"\n", "PROFILE", p.profile.value)

	data += ")\n"

	fmt.Println(data)
	os.WriteFile(envFilePath, []byte(data), 0666)
}
