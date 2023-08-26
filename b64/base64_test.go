package b64

import (
	"testing"
)

func Test_Encode(t *testing.T) {
	tests := []struct {
		Case     string
		Expected string
	}{
		{"", ""},
		{"fooba", "Zm9vYmE="},
		{"foobar", "Zm9vYmFy"},
		{"Zm9vYmFy", "Wm05dlltRnk="},
		{"Hello", "SGVsbG8="},
		{"World", "V29ybGQ="},
		{"Hello World", "SGVsbG8gV29ybGQ="},
	}

	for _, test := range tests {
		got, err := Encode(test.Case)
		if err != nil {
			t.Error("Expected ", test.Expected, ", got ", err)
		}
		if got != test.Expected {
			t.Error("Expected ", test.Expected, ", got ", got)
		}
	}
}

func Test_Decode(t *testing.T) {
	tests := []struct {
		Case     string
		Expected string
	}{
		{"", ""},
		{"VGhpcyBpcyBhIGJhc2U2NCBlbmNvZGVkIHN0cmluZyE=", "This is a base64 encoded string!"},
		{"Zm9vYmE=", "fooba"},
		{"Wm05dlltRnk=", "Zm9vYmFy"},
		{"SGVsbG8=", "Hello"},
		{"V29ybGQ=", "World"},
		{"SGVsbG8gV29ybGQ=", "Hello World"},
	}

	for _, test := range tests {
		got, err := Decode(test.Case)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got != test.Expected {
			t.Errorf("Case: %s failed, expected %v, got %v", test.Case, test.Expected, got)
		}
	}
}
