package b64

import "testing"

func Test_GetChar(t *testing.T) {
	tests := []struct {
		Case     rune
		Expected int
	}{
		{'A', 65},
		{'a', 97},
		{'0', 48},
		{'/', 47},
	}

	for _, test := range tests {
		if toCharCode(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", toCharCode(test.Case))
		}
	}
}

func Test_ToBinary(t *testing.T) {
	tests := []struct {
		Case     int
		Expected string
	}{
		{65, "1000001"},
		{97, "1100001"},
		{48, "110000"},
		{47, "101111"},
	}

	for _, test := range tests {
		if toBinary(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", toBinary(test.Case))
		}
	}
}
