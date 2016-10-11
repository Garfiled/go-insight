package main

import (
	"testing"
)

func BenchmarkEscape1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		i := 0
		pp := new(*int)
		*pp = &i
		_ = pp
	}
}
