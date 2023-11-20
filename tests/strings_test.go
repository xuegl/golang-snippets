package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacer(t *testing.T) {
	r := strings.NewReplacer("#", "o")
	broken := "G#phers"
	fixed := r.Replace(broken)
	assert.Equal(t, "Gophers", fixed)
}
