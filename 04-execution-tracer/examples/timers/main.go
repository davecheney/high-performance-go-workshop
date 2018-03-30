package main

import (
	"time"

	"github.com/pkg/profile"
)

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func main() {
	const N = 10
	defer profile.Start(profile.TraceProfile).Stop()
	for i := 0; i < N; i++ {
		c := timer(1 * time.Second)
		<-c
	}
}
