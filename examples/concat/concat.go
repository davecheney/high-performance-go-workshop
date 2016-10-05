// +build none

package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

// START OMIT

var request struct{ ID string }
var client net.Listener

func A() string {
	return request.ID + " " + client.Address().String() + " " + time.Now().String()
}

func B() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %v %v", request.ID, client.Addr(), time.Now())
	return b.String()
}

func C() string {
	b := make([]byte, 0, 40) // guess
	b = append(b, request.ID...)
	b = append(b, ' ')
	b = append(b, addr.String()...)
	b = append(b, ' ')
	b = time.Now().AppendFormat(b, "2006-01-02 15:04:05.999999999 -0700 MST")
	return string(b)
}

// END OMIT
