// +build none

package main

// tag::max[]
func F() {
	const a, b = 100, 20
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	if result == b {
		panic(b)
	}
}

// end::max[]

func main() {
	F()
}
