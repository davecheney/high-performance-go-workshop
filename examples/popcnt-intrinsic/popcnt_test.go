package main

import (
	"math/bits"
	"testing"
)

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

var Result uint64

func BenchmarkPopcnt(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = popcnt(uint64(i))
	}
	Result = r
}

// tag::benchmark[]
func BenchmarkMathBitsPopcnt(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = bits.OnesCount64(uint64(i))
	}
	Result = uint64(r)
}

// end::benchmark[]
