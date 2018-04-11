package main

import (
	"sync"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MutexProfile).Stop()

	var mu sync.Mutex
	var items = make(map[int]struct{})

	const N = 1000 * 1000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
		}(i)
	}
	wg.Wait()
}
