package p

import (
	"encoding/binary"
	"io"
)

// START OMIT
type Conn struct {
	r  io.ReadCloser
	ch chan uint32
}

func (c *Conn) Loop() {
	defer r.Close()
	var buf [512]byte
	for {
		b := buf[:] // create slice of buf
		n, err := c.r.Read(b)

		for b = b[:n]; len(b) != 0; b = b[4:] {
			ch <- binary.BigEndian.Uint32(b)
		}

		if err != nil {
			return
		}
	}
}

// END OMIT
