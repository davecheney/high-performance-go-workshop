package main

import (
	"sync/atomic"
)

type counter uint64

func (c *counter) get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
func (c *counter) inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}
func (c *counter) reset() uint64 {
	return atomic.SwapUint64((*uint64)(c), 0)
}

var c counter

func f() uint64 {
	c.inc()
	c.get()
	return c.reset()
}

func main() {
	f()
}
