package strings

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	s, err := Rand(32)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
}
