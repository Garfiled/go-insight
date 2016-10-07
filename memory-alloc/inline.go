package main

func main() {
	foo()
}

func foo() *int {
	a := 100
	return &a
}
