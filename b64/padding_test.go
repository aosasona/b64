package b64

import "testing"

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

func Test_PadOutput(t *testing.T) {
	// valid base64 strings are not used here
	tests := []struct {
		Case     string
		Expected string
	}{
		{"MQs", "MQs="},
		{"Kw", "Kw=="},
		{"SQkPA", "SQkPA==="},
	}

	for _, test := range tests {
		if padOutput(test.Case) != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", padOutput(test.Case))
		}
	}
}
