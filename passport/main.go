package main

import (
//	"github.com/fuzxxl/nfc/dev/nfc"
)

var desParity [256]byte

func parity(b byte) byte {
	var p byte
	for p = 0; b != 0; b >>= 1 {
		p += b & 1
	}
	return p & 1
}

func init() {
	var k int
	for k = 0; k < 256; k++ {
		n := (byte)(k & 0xfe)
		if parity(n) == 0 {
			n ^= 1
		}
		desParity[k] = n
	}
}
