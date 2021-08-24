package buildin

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//Build is struct base of constructor commands.
type Build struct {
	OS   string
	Arch string
	help string
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

//usage return of help for commands.
func usage() string {
	return `Simple build all plataforms bynaries write in go.
USAGE: buildin -os [OS] -arch [ARCH]
ARGS:
	-h      help commands.
	-os    	target operating system.
			eg. windows,linux,mac...
	-arch   destination architecture.
			eg. amd64, arm, ...
`
}

//parseArgs execute of parse the args from comand line.
func (b *Build) parseArgs() error {
	os := flag.String("os", runtime.GOOS, "target operating system.")
	arch := flag.String("arch", runtime.GOARCH, "destination architecture.")
	help := flag.String("h", usage(), "help commands.")

	flag.Parse()
	if len(*os) == 0 || len(*arch) == 0 {
		return ErrNotParse
	}
	b.OS = *os
	b.Arch = *arch
	b.help = *help
	return nil
}

//createDir is responsable by create dir build for deposit executables.
func (b *Build) createDir() error {
	if _, err := os.Stat("dist"); os.IsNotExist(err) {
		err = os.Mkdir("dist", 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

//createParams return params for cmd command.
func (b *Build) createParams() string {
	var output string
	osAndArch := fmt.Sprintf("GOOS=%v GOARCH=%v", b.OS, b.Arch)
	input := "main.go"
	if b.OS == "windows" {
		output = filepath.Join("dist", fmt.Sprintf("%v_%v.exe", b.OS, b.Arch))
	} else {
		output = filepath.Join("dist", fmt.Sprintf("%v_%v", b.OS, b.Arch))
	}
	cmpParams := fmt.Sprintf("%v go build -o %v %v", osAndArch, output, input)
	fmt.Printf("Build -> %v/%v.\n", b.OS, b.Arch)
	return cmpParams
}

//start create binary of os and arch.
func (b *Build) start() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("start", b.createParams())
	} else {
		cmd = exec.Command("bash", "-c", b.createParams())
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

//Run executing of search by initial file from actual dir and build program.
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

	if len(os.Args) < 4 || os.Args[1] == "-h" {
		fmt.Println(usage())
		os.Exit(0)
	}

	err = b.createDir()
	if err != nil {
		return err
	}

	err = b.parseArgs()
	if err != nil {
		return err
	}

	err = b.start()
	if err != nil {
		return err
	}

	return nil
}
