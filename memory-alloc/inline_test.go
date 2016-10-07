package main

import (
	"testing"
)

func BenchmarkInline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func foo() *int {
	a := 100
	return &a
}
