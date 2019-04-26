package main

import "testing"

var Result int

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()
	var r int
	for i := 0; i < b.N; i++ {
		r = Sum()
	}
	Result = r
}
