package main

import "testing"

// tag::bench[]
var X [1 << 15]struct {
	val int
	_   [4096]byte
}

var Result int

func BenchmarkRange(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		for _, x := range X {
			r += x.val
		}
	}
	Result = r
}

func BenchmarkFor(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(X); i++ {
			x := &X[i]
			r += x.val
		}
	}
	Result = r
}

// end::bench[]
