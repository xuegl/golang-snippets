package tests

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGlob(t *testing.T) {
	matches, err := filepath.Glob("*.go")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(matches)
}
