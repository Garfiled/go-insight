package main

import (
	"testing"
)

func BenchmarkClosureEscape1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var y int
		func(p int, x int) {
			p = x
		}(y, 42)
	}
}
func BenchmarkClosureEscape2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var y int
		func(p *int, x int) {
			*p = x
		}(&y, 42)
	}
}

func BenchmarkClosureEscape3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var y int
		func(x int) {
			y = x
		}(42)
	}
}

func BenchmarkClosureEscape4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := 100
		y := &a
		func(x int) {
			*y = x
		}(42)
	}
}
