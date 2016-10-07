// map_test.go
// go test -bench="." map_test.go -num 8
// go test -bench="." map_test.go -num 9

package main

import (
	"flag"
	"testing"
)

var num int

func init() {
	flag.IntVar(&num, "num", 1, "")
	flag.Parse()
}

func BenchmarkMakeMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < num; j++ {
			m[j] = j * 100
		}
	}
}
