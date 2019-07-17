package main

import "testing"

type T struct {
	padding [1 << 20]byte
	buffer  [256]byte
}

var Result byte

var t = new(T)

func BenchmarkNilCheckNotElided(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = t.buffer[1]
		Result = t.buffer[2]
		Result = t.buffer[4]
		Result = t.buffer[8]
		Result = t.buffer[16]
		Result = t.buffer[32]
		Result = t.buffer[64]
	}
}

func BenchmarkNilCheckElided(b *testing.B) {
	_ = t.buffer
	for n := 0; n < b.N; n++ {
		Result = t.buffer[1]
		Result = t.buffer[2]
		Result = t.buffer[4]
		Result = t.buffer[8]
		Result = t.buffer[16]
		Result = t.buffer[32]
		Result = t.buffer[64]

	}
}
