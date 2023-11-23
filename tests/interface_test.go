package tests

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	println("nil interface:")
	printNilInterface()
	println("empty interface:")
	printEmptyInterface()
	println("non-empty interface:")
	printNonEmptyInterface()
}

func printNilInterface() {
	var i interface{}
	var e error
	println(i)
	println(e)
	println("i == nil: ", i == nil)
	println("e == nil: ", e == nil)
	println("i == e: ", i == e)
}

func printEmptyInterface() {
	var eif1 interface{}
	var eif2 interface{}
	n, m := 17, 18

	eif1 = n
	eif2 = m
	println("eif1: ", eif1)
	println("eif2: ", eif2)
	println("eif1 == eif2: ", eif1 == eif2)

	eif2 = 17
	println("eif1: ", eif1)
	println("eif2: ", eif2)
	println("eif1 == eif2: ", eif1 == eif2)

	eif2 = int64(17)
	println("eif1: ", eif1)
	println("eif2: ", eif2)
	println("eif1 == eif2: ", eif1 == eif2)
}

type IT struct {
	n int
	s string
}

func (it IT) Error() string {
	return "IT"
}

func printNonEmptyInterface() {
	var err1 error
	var err2 error
	err1 = (*IT)(nil)
	println("err1: ", err1)
	println("err1 == nil: ", err1 == nil)

	err1 = IT{5, "IT"}
	err2 = IT{6, "IT"}
	println("err1: ", err1)
	println("err2: ", err2)
	println("err1 == err2: ", err1 == err2)

	err2 = fmt.Errorf("%d", 5)
	println("err1: ", err1)
	println("err2: ", err2)
	println("err1 == err2: ", err1 == err2)
}
