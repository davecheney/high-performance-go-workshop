package main

import "testing"

type T struct {
	padding [8192]byte
	buffer  [256]byte
}

var Result byte

var t = new(T)

func BenchmarkNilCheckNotElided(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Result = t.buffer[10]
	}
}

func BenchmarkNilCheckElided(b *testing.B) {
	_ = t.buffer
	for n := 0; n < b.N; n++ {
		Result = t.buffer[10]
	}
}
