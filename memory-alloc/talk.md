# About Go Memory Allocation

liuguixiang

BeiJing 2016-10-7

## Env

* Go1.7

* OS X 

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
### Test && Compile
```
$ go test -bench="." array_test.go
$ go tool compile -S array.go
```



