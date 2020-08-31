package config

import (
	"testing"
)

func TestLength_Truncate(t *testing.T) {
	cases := []struct {
		label       string
		length      Length
		number      int
		expectedMod int
	}{
		{"Length 1", Length1, 433494437, 7},
		{"Length 2", Length2, 433494437, 37},
		{"Length 3", Length3, 433494437, 437},
		{"Length 4", Length4, 433494437, 4437},
		{"Length 5", Length5, 433494437, 94437},
		{"Length 6", Length6, 433494437, 494437},
		{"Length 7", Length7, 433494437, 3494437},
		{"Length 8", Length8, 433494437, 33494437},
		{"Length 10", Length(10), 433494437, 433494437},
	}

	for _, c := range cases {
		modulus := c.length.Truncate(c.number)

		if c.expectedMod != modulus {
			t.Errorf("case %s: wrong modulus\nexpected: %d\n  actual: %d", c.label, c.expectedMod, modulus)
			t.FailNow()
		}
	}
}

func TestLength_LeftPad(t *testing.T) {
	cases := []struct {
		label    string
		length   Length
		number   int
		expected string
	}{
		{"Length 1", Length1, 1, "1"},
		{"Length 2", Length2, 1, "01"},
		{"Length 3", Length3, 1, "001"},
		{"Length 4", Length4, 1, "0001"},
		{"Length 5", Length5, 1, "00001"},
		{"Length 6", Length6, 1, "000001"},
		{"Length 7", Length7, 1, "0000001"},
		{"Length 8", Length8, 1, "00000001"},
		{"Length 10", Length(10), 1, "0000000001"},
		{"Long number", Length2, 433494437, "433494437"},
	}

	for _, c := range cases {
		padded := c.length.LeftPad(c.number)

		if c.expected != padded {
			t.Errorf("case %s: wrong modulus\nexpected: %s\n  actual: %s", c.label, c.expected, padded)
			t.FailNow()
		}
	}
}
