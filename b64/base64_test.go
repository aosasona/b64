package b64

import (
	"reflect"
	"testing"
)

func Test_ToBinary(t *testing.T) {
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

func Test_PadLeft8bitBinary(t *testing.T) {
	tests := []struct {
		Case     string
		Expected string
	}{
		{"1000001", "01000001"},
		{"1111001", "01111001"},
		{"1101111", "01101111"},
	}

	for _, test := range tests {
		if padLeft8bitBinary(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", padLeft8bitBinary(test.Case))
		}
	}
}

func Test_PadRight(t *testing.T) {
	tests := []struct {
		Case     string
		Expected string
	}{
		{"101", "101000"},
		{"10110", "101100"},
		{"101101", "101101"},
	}

	for _, test := range tests {
		if padRight(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", padRight(test.Case))
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
