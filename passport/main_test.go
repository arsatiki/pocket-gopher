package main

import (
	"testing"
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
