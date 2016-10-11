# About Go Memory Allocation

liuguixiang

BeiJing 2016-10-7

## Environment

* Go 1.7

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
Constant propagation happens in SSA, which happens after escape analysis
and I have proposed a log issue, Go1.8 Maybe fix it .
https://github.com/golang/go/issues/17275

```go
// slicemake_test.go
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
// slicemake.go
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
## Map
```go
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
```

```go
// map.go
package main

func main() {
	m := make(map[int]int)
	for j := 0; j < 8; j++ {
		m[j] = j * 100
	}
}
```
## For Statement Escape
```go
// forescape_test.go
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
```
```go
// forescape.go
package main

func main() {
	var a []int

	for i := 0; i < 1; i++ {
		// a := make([]int, 0, 2)
		a = make([]int, 0, 2)
	}
	_ = a
}
```
## Label Statement Escape
```go
// lableescape_test.go
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
```
```go
// labelescape.go
package main

func main() {
	var a []int

Label1:
	a = make([]int, 0, 2)
	if len(a) > 2 {
		goto Label1
	}
	_ = a
}
```
## Interface
```go
// interface_test.go
package main

import (
	"testing"
)

const Width, Height = 640, 480

type CenterInf interface {
	Center()
}

type Cursor struct {
	X, Y int
}

func (c *Cursor) Center() {
	c.X += Width / 2
	c.Y += Height / 2
}

func BenchmarkInterfaceNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var a interface{}
		_ = a
	}
}
func BenchmarkInterfaceValue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := interface{}(100)
		_ = a
	}
}

// 虽然是new但是对象还是在栈上分配的
// 对象没有发生逃逸，而且是小对象
func BenchmarkNewObject(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := new(Cursor)
		c.Center()
	}
}

// 变量转interface{} 本身不会发生逃逸
func BenchmarkNewObjectInterface(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := new(Cursor)
		t := CenterInf(c)
		_ = t
	}
}

// 调用接口变量的方法会导致变量逃逸
func BenchmarkNewObjectEscape(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := new(Cursor)
		CenterInf(c).Center()
	}
}
```
```go
// interface.go
package main

const Width, Height = 640, 480

type CenterInf interface {
	Center()
}

type Cursor struct {
	X, Y int
}

func (c *Cursor) Center() {
	c.X += Width / 2
	c.Y += Height / 2
}

func main() {
	c := new(Cursor)
	c.Center()
}

// func main() {
// 	c := new(Cursor)
// 	CenterInf(c).Center()
// }
```
## Func Inline
```go
package main

import (
	"testing"
)

func BenchmarkInline(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func foo() *int {
	a := 100
	return &a
}
```
```go
// inline.go
package main

func main() {
	foo()
}

func foo() *int {
	a := 100
	return &a
}
```
## Fmt
```go
// fmtalloc_test.go
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
```
