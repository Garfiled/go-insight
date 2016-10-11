package main

import "fmt"

func main() {
	println("main>>>")
	s := []int{100}
	i := interface{}(s)
	println("debug")
	fmt.Println(i)
	println("end")
}
