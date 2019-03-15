// +build none

package main

import "testing"

// tag::benchmark[]
func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// optimised away
	}
}

// end::benchmark[]
