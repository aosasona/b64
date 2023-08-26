package b64

import "testing"

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
