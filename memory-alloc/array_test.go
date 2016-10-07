package main

import "testing"

// The threshold for stack allocation is 10MB
func BenchmarkEaArray1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var arr [1024 * 1024 * 10]byte
		_ = arr
	}
}
func BenchmarkEaArray2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var arr [1024*1024*10 + 1]byte
		_ = arr
	}
}
