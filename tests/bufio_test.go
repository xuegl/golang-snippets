package tests

import (
	"errors"
	"fmt"
	"testing"
)

func TestBufio_New(t *testing.T) {
	err := errors.New("test")
	fmt.Printf("err: %v\n", err)
}
