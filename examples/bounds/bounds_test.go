package main

import "testing"

// tag::inorder[]
var v = make([]int, 9)

var A, B, C, D, E, F, G, H, I int

func BenchmarkBoundsCheckInOrder(b *testing.B) {
	var a, _b, c, d, e, f, g, h, i int
	for n := 0; n < b.N; n++ {
		a = v[0]
		_b = v[1]
		c = v[2]
		d = v[3]
		e = v[4]
		f = v[5]
		g = v[6]
		h = v[7]
		i = v[8]
	}
	A, B, C, D, E, F, G, H, I = a, _b, c, d, e, f, g, h, i
}

// end::inorder[]

// tag::outoforder[]
func BenchmarkBoundsCheckOutOfOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		I = v[8]
		A = v[0]
		B = v[1]
		C = v[2]
		D = v[3]
		E = v[4]
		F = v[5]
		G = v[6]
		H = v[7]
	}
}

// end::outoforder[]
