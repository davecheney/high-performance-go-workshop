package fib

import "testing"

func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// tag::wrong[]
func BenchmarkFibWrong(b *testing.B) {
	Fib(b.N)
}

// end::wrong[]

// tag::wrong2[]
func BenchmarkFibWrong2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
}

// end::wrong2[]

// END OMIT
