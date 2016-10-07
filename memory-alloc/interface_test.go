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
