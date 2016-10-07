package main

import (
	"testing"
)

func BenchmarkSliceAppend(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p := make([]int, 0, 1)
		p = append(p, 1)
		p = append(p, 2)
		p = append(p, 3)
		p = append(p, 4)
		p = append(p, 5)
		p = append(p, 6)
		p = append(p, 7)
	}
}
