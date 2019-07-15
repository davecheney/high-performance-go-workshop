package main

import (
	"sync"
	"testing"
)

// tag::mutex[]
type AtomicVariable struct {
	mu  sync.Mutex
	val uint64
}

func (av *AtomicVariable) Inc() {
	av.mu.Lock()
	av.val++
	av.mu.Unlock()
}

func BenchmarkInc(b *testing.B) {
	var av AtomicVariable

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			av.Inc()
		}
	})
}

// end::mutex[]
