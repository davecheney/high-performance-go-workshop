package main

import (
	"testing"
	"time"
)

var Result int

func BenchmarkStartStop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		Result++
		b.StartTimer()
		Result += int(time.Now().Unix())
	}
}
