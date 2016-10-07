package main

// make是对于for作用域的，赋值操作
func main() {
	var a []int

	for i := 0; i < 1; i++ {
		// a := make([]int, 0, 2)
		a = make([]int, 0, 2)
	}
	_ = a
}
