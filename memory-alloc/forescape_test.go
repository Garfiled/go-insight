package main

import (
	"testing"
)

func BenchmarkForEscape1(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1; i++ {
			a := make([]int, 0, 2)
			_ = a
		}
	}
}

func BenchmarkForEscape2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var a []int
		for i := 0; i < 1; i++ {
			a = make([]int, 0, 2)
		}
		_ = a
	}
}
