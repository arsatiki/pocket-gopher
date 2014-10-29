package main

import (
	//	"github.com/fuzxxl/nfc/dev/nfc"
	"crypto/sha1"
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

func key(seed []byte, suffix []byte) []byte {
	buf := make([]byte, 20, 20)
	copy(buf, seed)
	copy(buf[16:], suffix)	

	h0 := sha1.Sum(buf)
	h := h0[:16]

	for i, v := range h {
		h[i] = desParity[v]
	}
	return h
}

func GetBACKeys(passport, dob, expiry string) (enc, mac []byte) {
	seed := sha1.Sum([]byte(passport + dob + expiry))
	enc = key(seed[:16], []byte{0, 0, 0, 1})
	mac = key(seed[:16], []byte{0, 0, 0, 2})

	return enc, mac
}
