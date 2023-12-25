package tests

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafePoint(t *testing.T) {
	type user struct {
		name string
		age  int
		pets []string
	}

	var u user
	fmt.Printf("%v+\n", u)

	uNamePtr := (*string)(unsafe.Pointer(&u))
	*uNamePtr = "snow"
	fmt.Printf("%+v\n", u)

	uAgePtr := (*int)(unsafe.Add(unsafe.Pointer(&u), unsafe.Offsetof(u.age)))
	*uAgePtr = 18
	fmt.Printf("%+v\n", u)

	u.pets = []string{"cat", "dog", "fish"}
	fmt.Printf("%+v\n", u)

	// Now we want to get a pointer to the second slice element and make
	// a change to it. We use the new unsafe func here called 'SliceData'.
	// This will return a pointer to the underlying array of the argument
	// slice. Now that we have a pointer to the array, we can add the size
	// of one string to the pointer to get the address of the second element.
	// This means you could say 2*unsafe.Sizeof("") to get to the third
	// element in this example if that is helpful for visualizing.
	secondAnimal := (*string)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(u.pets)), unsafe.Sizeof("")))
	*secondAnimal = "rabbit"
	fmt.Printf("%+v\n", u)

	// Iterate each element in a slice
	p := unsafe.Pointer(unsafe.SliceData(u.pets))
	es := unsafe.Sizeof(u.pets[0])
	for i := 0; i < len(u.pets); i++ {
		fmt.Println(*(*string)(unsafe.Add(p, uintptr(i)*es)))
	}
}

func TestUnsafeString(t *testing.T) {
	buf := []byte{'H', 'e', 'l', 'l', 'o'}
	println(unsafe.String(unsafe.SliceData(buf), len(buf)))
}

func TestUnsafeStringData(t *testing.T) {
	s := "Hello"
	buf := unsafe.Slice(unsafe.StringData(s), len(s))
	for _, b := range buf {
		println(b)
	}
}

func TestUnsafe_Finalizer(t *testing.T) {
	type Foo struct {
		name string
	}

	finalizer := func(f *Foo) {
		fmt.Printf("Foo[%s] is recycled\n", f.name)
	}

	newFoo := func(name string) *Foo {
		foo := &Foo{
			name: name,
		}
		runtime.SetFinalizer(foo, finalizer)
		return foo
	}

	allocLargeObj := func() *[1000000]int64 {
		a := [1000000]int64{}
		return &a
	}

	p1 := uintptr(unsafe.Pointer(newFoo("FooRefByUnitptr")))
	p2 := unsafe.Pointer(newFoo("FooRefByPointer"))

	for i := 0; i < 5; i++ {
		allocLargeObj()
		q1 := (*Foo)(unsafe.Pointer(p1))
		fmt.Printf("object ref by uintptr: %+v\n", *q1)

		q2 := (*Foo)(p2)
		fmt.Printf("object ref by pointer: %+v\n", *q2)

		runtime.GC()
		time.Sleep(1 * time.Second)
	}
}
