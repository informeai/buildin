package buildin

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

//Build is struct base of constructor commands.
type Build struct {
	OS        string
	Arch      string
	inputDir  string
	outputDir string
	all       bool
}

//errors
var (
	ErrEmptyValue = errors.New("empty value not permited")
	ErrNotFound   = errors.New("value not found")
	ErrNotParse   = errors.New("not parsed args")
	ErrOs         = errors.New("os not accepted")
	ErrArch       = errors.New("arch not accepted")
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
func NewBuild() *Build {
	return &Build{OS: runtime.GOOS, Arch: runtime.GOARCH}
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
	return false, ErrOs
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
	return false, ErrArch
}

//parseArgs execute of parse the args from comand line.
func (b *Build) parseArgs() error {
	actualDir, err := os.Getwd()
	if err != nil {
		return err
	}
	out := filepath.Join(actualDir, "build")
	os := flag.String("os", runtime.GOOS, "target operating system.")
	arch := flag.String("arch", runtime.GOARCH, "destination architecture.")
	input := flag.String("i", actualDir, "current directory for build.")
	output := flag.String("o", out, "destination directory for build.")
	all := flag.Bool("all", false, "build for everyone.")

	flag.Parse()
	if len(*os) == 0 || len(*arch) == 0 {
		return ErrNotParse
	}
	b.OS = *os
	b.Arch = *arch
	b.all = *all
	b.inputDir = *input
	b.outputDir = *output
	return nil
}

//Run executing of search by initial file in inputDir and build program.
//eg. In programs of go the search (main.go) by default.
func (b *Build) Run() error {
	t, err := verifyOs(b.OS)
	if t == false && err != nil {
		return err
	}
	t, err = verifyArch(b.Arch)
	if t == false && err != nil {
		return err
	}

	return nil
}
