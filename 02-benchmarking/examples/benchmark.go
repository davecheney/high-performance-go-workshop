package q

import "testing"

func BenchmarkRead(b *testing.B) {
	b.ReportAllocs() // HL
	for n := 0; n < b.N; n++ {
		// function under test
	}
}
