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
