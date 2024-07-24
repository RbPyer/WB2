package tests

import (
	"dev02/unpack"
	"errors"
	"testing"
)

type TestCase struct {
	received string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {

	tests := []TestCase{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", unpack.InvalidStringErr},
		{"", "", nil},
	}

	for _, test := range tests {
		result, err := unpack.StringUnpack(test.received)
		if result != test.expected || !errors.Is(err, test.err) {
			t.Errorf("UnpackString(%q) = %q, %v; want %q, %v",
				test.received, result, err, test.expected, test.err)
		}
	}
}
