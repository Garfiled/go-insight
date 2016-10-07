package main

import (
	"fmt"
	"testing"
)

func BenchmarkFmtEscap1(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		fmt.Sprint(100)
	}
}
func BenchmarkFmtEscape2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		fmt.Sprint(struct {
			Field int
			Age   int
		}{10, 27})
	}
}
func BenchmarkFmtEscap3(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		a := interface{}(100)
		fmt.Sprint(a)
	}
}
