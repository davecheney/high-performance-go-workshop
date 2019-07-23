package main

// tag::max[]
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func F() {
	var a, b = 100, 20
	if Max(a, b) == b {
		panic(b)
	}
}

// end::max[]

func main() {
	F()
}
