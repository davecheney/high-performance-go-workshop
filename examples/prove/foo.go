package main

func foo(x int) bool {
	if x > 5 { // <1>
		if x > 3 { // <2>
			return true
		}
		panic("x less than 3")
	}
	return false
}

func main() {
	foo(-1)
}
