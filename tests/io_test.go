package tests

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestIO_Copy(t *testing.T) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, strings.NewReader("Hello World"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", buf.String())

	buf.Reset()
	_, err = io.Copy(&buf, bytes.NewReader([]byte("Hello World")))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", buf.String())
}
