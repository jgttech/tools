package str

func NewSemVer(version string) semVer {
	var v semVer

	v.Load(version)

	return v
}
