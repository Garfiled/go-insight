package main

import (
	"testing"
)

func BenchmarkLabelEscape(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var a []int
	Label1:
		a = make([]int, 0, 2)
		if len(a) > 2 {
			goto Label1
		}
		_ = a
	}
}
