package main

import (
	"fmt"
	"unsafe"
)

// tag::struct[]
type S struct {
	a bool
	b float64
	c int32
}

// end::struct[]

// tag::padding[]
type S struct {
	a bool
	_ [7]byte // padding <1>
	b float64
	c int32
	_ [4]byte // padding <2>
}

// end::padding[]

func main() {
	// tag::sizeof[]
	var s S
	fmt.Println(unsafe.Sizeof(s)) // <1>
	// end::sizeof[]
}
