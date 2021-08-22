package buildin

import (
	"errors"
)

//Build is struct base of constructor commands.
type Build struct {
	OS   string
	Arch string
}

//errors
var (
	ErrEmptyValue = errors.New("empty value not permited")
	ErrNotFound   = errors.New("value not found")
)

//goos
var GOOS = []string{
	"aix",
	"android",
	"darwin",
	"dragonfly",
	"freebsd",
	"hurd",
	"illumos",
	"js",
	"linux",
	"nacl",
	"netbsd",
	"openbsd",
	"plan9",
	"solaris",
	"windows",
	"zos",
}

//goarch
var GOARCH = []string{
	"386",
	"amd64",
	"amd64p32",
	"arm",
	"armbe",
	"arm64",
	"arm64be",
	"ppc64",
	"ppc64le",
	"mips",
	"mipsle",
	"mips64",
	"mips64le",
	"mips64p32",
	"mips64p32le",
	"ppc",
	"riscv",
	"riscv64",
	"s390",
	"s390x",
	"sparc",
	"sparc64",
	"wasm",
}

//NewBuild return new instance of Build.
func NewBuild(os, arch string) *Build {
	return &Build{OS: os, Arch: arch}
}

//verifyOs return true if os exists in list of GOOS.
func verifyOs(os string) (bool, error) {
	if len(os) == 0 {
		return false, ErrEmptyValue
	}
	for _, o := range GOOS {
		if os == o {
			return true, nil
		}
	}
	return false, ErrNotFound
}

//verifyArch return true if arch exists in list of GOARCH.
func verifyArch(arch string) (bool, error) {
	if len(arch) == 0 {
		return false, ErrEmptyValue
	}
	for _, a := range GOARCH {
		if arch == a {
			return true, nil
		}
	}
	return false, ErrNotFound
}
