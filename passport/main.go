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



func key(seed [20]byte, suffix [4]byte) [16]byte {
	var res [16]byte

	buf := make([]byte, 20)
	copy(buf, seed[:16])
	copy(buf[16:], suffix[:])

	h := sha1.Sum(buf)
	copy(res[:], h[:16])

	for i, v := range res {
		res[i] = desParity[v]
	}
	return res
}

func GetBACKeys(passport, dob, expiry string) (enc, mac [16]byte) {
	seed := sha1.Sum([]byte(passport + dob + expiry))
	enc = key(seed, [4]byte{0, 0, 0, 1})
	mac = key(seed, [4]byte{0, 0, 0, 2})

	return enc, mac
}
