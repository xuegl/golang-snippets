package strings_test

import (
	"fmt"
	"github.com/snow/golang-snippets/strings"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		s, err := strings.Rand(10)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(s)
	}
}
