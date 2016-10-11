package main

func main() {
	var y int // BAD: y escapes
	func(p *int, x int) {
		*p = x
	}(&y, 42)
}

// func main() {
// 	var y int
// 	func(x int) {
// 		y = x
// 	}(42)
// }
//
// func main() {
// 	a := 100
// 	y := &a
// 	func(x int) {
// 		*y = x
// 	}(42)
// }
