package main

// BAD: i escapes
func main() {
	i := 0
	pp := new(*int)
	*pp = &i
	_ = pp
}
