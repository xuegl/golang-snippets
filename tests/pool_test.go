package tests

import (
	"bytes"
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func writeWithoutPool(s string) {
	buf := new(bytes.Buffer)
	buf.WriteString(s)
}

func writeWithPool(s string) *bytes.Buffer {
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString(s)
	pool.Put(buf)
	return buf
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeWithoutPool("hello")
	}
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeWithPool("hello")
	}
}
