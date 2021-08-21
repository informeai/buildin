package buildin

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewBuild
func TestNewBuild(t *testing.T) {
	b := NewBuild("darwin", "amd64")
	if b == nil {
		t.Errorf("TestNewBuild(): got -> %v, want: Build{}", b)
	}
	log.Println(b)
}
