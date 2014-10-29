package main

import (
	"testing"
	"bytes"
)

func TestParity(t *testing.T) {
	if parity(1) != 1 {
		t.Error("Expected 1, got 0")
	}
	if parity(2) != 1 {
		t.Error("Expected 1, got 0")
	}
	if parity(0xf) != 0 {
		t.Error("Expected 0, got 1")
	}
}

func TestDesparity(t *testing.T) {
	var data = []struct {
		input, expected byte
	}{
		{0x1, 0x1},
		{0x9, 0x8},
	}
	for _, item := range data {
		output := desParity[item.input]
		if output != item.expected {
			t.Error("Expected", item.expected, "got", output)
		}

	}
}

func TestKeys(t *testing.T) {
	var expected_enc = []byte{
		0xab, 0x94, 0xfd, 0xec, 0xf2, 0x67, 0x4f, 0xdf,
		0xb9, 0xb3, 0x91, 0xf8, 0x5d, 0x7f, 0x76, 0xf2,
	}
	var expected_mac = []byte{
		0x79, 0x62, 0xd9, 0xec, 0xe0, 0x3d, 0x1a, 0xcd,
		0x4c, 0x76, 0x08, 0x9d, 0xce, 0x13, 0x15, 0x43,
	}
	var (
		passport = "L898902C<3"
		dob      = "6908061"
		expiry   = "9406236"
	)
	enc, mac := GetBACKeys(passport, dob, expiry)
	if !bytes.Equal(enc, expected_enc) {
		t.Fail()
	}
	if !bytes.Equal(mac, expected_mac) {
		t.Fail()
	}

}
