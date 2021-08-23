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
