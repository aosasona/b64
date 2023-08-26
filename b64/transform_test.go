package b64

import (
	"reflect"
	"testing"
)

func Test_AsciiToBin(t *testing.T) {
	tests := []struct {
		Case     rune
		Expected string
	}{
		{65, "01000001"},
		{97, "01100001"},
		{48, "00110000"},
	}

	for _, test := range tests {
		if asciiToBin(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", asciiToBin(test.Case))
		}
	}
}

func Test_BinToAscii(t *testing.T) {
	tests := []struct {
		Case     string
		Expected rune
	}{
		{"01000001", 65},
		{"01100001", 97},
		{"00110000", 48},
	}

	for _, test := range tests {
		r, err := binToAscii(test.Case)
		if err != nil {
			t.Error("Expected ", test.Expected, ", got ", err)
		}

		if r != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", r)
		}
	}
}

func Test_Chunk(t *testing.T) {
	tests := []struct {
		Case     string
		Expected []string
	}{
		{"101000", []string{"101000"}},
		{"1101011110", []string{"110101", "111000"}},
		{"001011010", []string{"001011", "010000"}},
		{"000101101000101", []string{"000101", "101000", "101000"}},
	}

	for _, test := range tests {
		if len(chunk(test.Case)) != len(test.Expected) {
			t.Error("Expected ", len(test.Expected), ", got ", len(chunk(test.Case)))
		}

		if !reflect.DeepEqual(chunk(test.Case), test.Expected) {
			t.Error("Expected ", test.Expected, ", got ", chunk(test.Case))
		}
	}
}

func Test_Chunk8bit(t *testing.T) {
	tests := []struct {
		Case     string
		Expected []string
	}{
		{"101000", []string{"00101000"}},
		{"1101011110", []string{"11010111", "00000010"}},
		{"000101101000101", []string{"00010110", "01000101"}},
		{"100110100011101010000100", []string{"10011010", "00111010", "10000100"}},
	}

	for _, test := range tests {
		if len(chunk8bit(test.Case)) != len(test.Expected) {
			t.Error("Expected ", len(test.Expected), ", got ", len(chunk8bit(test.Case)))
		}

		if !reflect.DeepEqual(chunk8bit(test.Case), test.Expected) {
			t.Error("Expected ", test.Expected, ", got ", chunk8bit(test.Case))
		}
	}
}
