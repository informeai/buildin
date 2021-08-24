package buildin

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewBuild
func TestNewBuild(t *testing.T) {
	b := NewBuild()
	if b == nil {
		t.Errorf("TestNewBuild(): got -> %v, want: Build{}", b)
	}
	log.Println(b)
}

//go test -v -run ^TestVerifyOs
func TestVerifyOs(t *testing.T) {
	b, err := verifyOs("solaris")
	if err != nil {
		t.Errorf("TestVerifyOs(): got -> %v", err)
	}
	log.Println(b)
}

//go test -v -run ^TestVerifyArch
func TestVerifyArch(t *testing.T) {
	b, err := verifyArch("amd64")
	if err != nil {
		t.Errorf("TestVerifyArch(): got -> %v", err)
	}
	log.Println(b)
}

//go test -v -run ^TestParseArgs
func TestParseArgs(t *testing.T) {
	b := NewBuild()
	err := b.parseArgs()
	if err != nil {
		t.Errorf("TestParseArgs(): got -> %v", err)
	}
}

//go test -v -run ^TestCreateDir
func TestCreateDir(t *testing.T) {
	b := NewBuild()
	b.inputDir = "../"
	log.Println(b.outputDir)
	err := b.createDir()
	if err != nil {
		t.Errorf("TestCreateDir(): got -> %v", err)
	}
}

//go test -v -run ^TestRun
func TestRun(t *testing.T) {
	b := NewBuild()
	err := b.Run()
	if err != nil {
		t.Errorf("TestRun(): got -> %v", err)
	}
}
