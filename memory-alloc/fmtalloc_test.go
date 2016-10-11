package main

import (
	"fmt"
	_ "net/http/pprof"
	"testing"
)

func BenchmarkFmt0(b *testing.B) {
	b.ReportAllocs()
	// var a interface{} = 100
	// 0 bytes/op    0 allocs/op
	// var a interface{} = []int{}
	// 0 bytes/op    0 allocs/op
	// var a interface{} = []int{100}
	// 8 bytes/op    1 allocs/op
	// var a interface{} = []int{100, 200}
	// 16 bytes/op    2 allocs/op
	var a interface{} = []interface{}{100}
	// 0 bytes/op    0 allocs/op
	// var a interface{} = [...]int{100, 200}
	// 0 bytes/op    0 allocs/op
	// var a interface{} = map[int]int{1: 100}
	// 144 bytes/op    4 allocs/op
	// var a interface{} = "hello world!"
	// 0 bytes/op    0 allocs/op
	// var a interface{} = []string{"hello", "world"}
	// 32 bytes/op    2 allocs/op
	// var a interface{} = []interface{}{"hello", "world"}
	// 0 bytes/op    0 allocs/op
	for i := 0; i < b.N; i++ {
		fmt.Println(a)
	}
}

// 8 Bytes/op   1 allocs/op
func _BenchmarkFmt1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Println(100)
	}
}

// 16 Bytes/op   1 allocs/op
func _BenchmarkFmt2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
}

// 32 Bytes/op   1 allocs/op
// request 24 bytes but the according to size_class alloc 32 bytes
func _BenchmarkFmt3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Println([]int{})
	}
}

// 48 Bytes/op   3 allocs/op
// 32 bytes for slice header
// 8 bytes for []int value
// 8 bytes for fmt
func _BenchmarkFmt4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Println([]int{100})
	}
}

// 64 Bytes/op   4 allocs/op
// 32 bytes for slice header
// 16 bytes for []int value
// 16 bytes for fmt
func _BenchmarkFmt5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Println([]int{100, 200})
	}
}

func _BenchmarkFmt6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := interface{}(100)
		if i < 0 {
			fmt.Printf("%d", a)
		}
	}
}
