package tests

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func TestReflect_NumMethods(t *testing.T) {
	// stdlib
	dumpMethods((*io.Reader)(nil))
	dumpMethods((*io.Writer)(nil))
	dumpMethods((*io.Closer)(nil))
	dumpMethods((*io.ReadWriter)(nil))
	dumpMethods((*io.WriteCloser)(nil))
	dumpMethods((*io.ReadWriteCloser)(nil))
	// user-defined
	var t1 T
	var pt1 *T
	dumpMethods(&t1)
	dumpMethods(&pt1)
	dumpMethods((*Interface)(nil))
}

func dumpMethods(i any) {
	t := reflect.TypeOf(i)
	et := t.Elem()
	fmt.Printf("%s's method set: \n", et)
	for i := 0; i < et.NumMethod(); i++ {
		fmt.Printf("\t- %s\n", et.Method(i).Name)
	}
}
