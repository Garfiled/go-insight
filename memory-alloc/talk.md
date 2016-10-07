# About Go Memory Allocation

liuguixiang

BeiJing 2016-10-7

## Env

* Go1.7

* OS X 

## Test && Compile && Escape Analysis 
```
$ go test -bench="." xxx_test.go
$ go tool compile -S xxx.go
$ go build -gcflags "-m" xxx.go
$ go build -gcflags "-m -m" xxx.go
$ go build -gcflags "-m -m -m -m" xxx.go
```
## Stack Allocation Limit
```go
// array_test.go
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
```

```go
// array.go
package main

func main() {
	var arr [1024 * 1024 * 10]byte
	_ = arr
}

// func main() {
// 	var arr [1024 * 1024 * 10+1]byte
// 	_ = arr
// }
```
## Slice Make
```go
// makeslice.go
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
```

```go
// makeslice.go
package main

func main() {
	s := make([]byte, 1024)
	_ = s
}

// func main() {
// 	cap := 1024
// 	s := make([]byte, cap)
// 	_ = s
// }

// func main() {
// 	s := make([]byte, 1024*64)
// 	_ = s
// }
```
## Slice Append
```go
// sliceappend_test.go
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
```
```go
// sliceappend.go
package main

func main() {
	p := make([]int, 0, 1)
	p = append(p, 1)
	p = append(p, 2)
	p = append(p, 3)
	p = append(p, 4)
	p = append(p, 5)
	p = append(p, 6)
	p = append(p, 7)
}
```
