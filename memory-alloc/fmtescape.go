package main

import (
	"fmt"
	. "unsafe"
)

/*
func main() {
	a := 100
	fmt.Println(a)
}
*/

/*
func main() {
	fmt.Sprint(100)
}
*/

func main() {
	println("debug>>>")
	a := interface{}(100)
	println("debug1>>>", &a, a)
	b := fmt.Sprint(a)
	println("debug2>>>", b, &b, Pointer(*((*uintptr)(Pointer(&b)))), *((*int)(Pointer(uintptr(Pointer(&b)) + 8))))
}
