package tests

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFile_ReadWrite(t *testing.T) {
	temp, err := os.CreateTemp("", "golang-snippets-")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = temp.Close()
	}()
	fmt.Printf("temp file: %s\n", temp.Name())
	_, err = temp.Write([]byte("Hello World"))
	if err != nil {
		t.Error(err)
	}
	_ = temp.Sync()

	buf := make([]byte, 20)
	_, _ = temp.Seek(0, 0)
	n, err := temp.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			println("EOF")
			return
		} else {
			t.Error(err)
		}
	}
	fmt.Printf("buf: %s\n", buf[:n])

	n, err = temp.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			println("EOF")
			return
		} else {
			t.Error(err)
		}
	}
	fmt.Printf("buf: %s\n", buf[:n])
}
