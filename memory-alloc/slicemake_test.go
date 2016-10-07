package main

import (
	"testing"
)

// s will be allocated in stack
func BenchmarkMakeSlice1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 0, 1024)
		_ = s
	}
}

// s will be allocated in heap
func BenchmarkMakeSlice2(b *testing.B) {
	b.ReportAllocs()
	cap := 1024
	for i := 0; i < b.N; i++ {
		s := make([]byte, cap, cap)
		_ = s
	}
}

// slice >= 64KB 对象被分配在了堆上面
func BenchmarkMakeSlice64KB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 0, 1024*64)
		_ = s
	}
}
