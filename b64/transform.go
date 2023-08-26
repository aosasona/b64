package b64

import (
	"fmt"
	"strconv"
)

func asciiToBin(code rune) string {
	return padLeft8bitBinary(strconv.FormatInt(int64(code), 2))
}

func binToAscii(bin string) (rune, error) {
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return rune(0), fmt.Errorf("Could not convert binary to ascii: %s", err.Error())
	}
	return rune(i), nil
}

// chunk and chunk8bit could be combined into one function but they are separated for clarity hre
// this is not written for performance at all, it is to express what is going on in a simple way
func chunk(str string) []string {
	var chunks []string
	str = padRight(str)

	for i := 0; i < len(str); i += 6 {
		end := i + 6
		if end > len(str) {
			end = len(str)
		}
		chunks = append(chunks, str[i:end])
	}

	return chunks
}

func chunk8bit(str string) []string {
	var chunks []string

	for i := 0; i < len(str); i += 8 {
		end := i + 8
		if end > len(str) {
			end = len(str)
		}

		bits := str[i:end]
		if len(bits) < 8 {
			bits = padLeft8bitBinary(bits)
		}

		chunks = append(chunks, bits)
	}

	return chunks
}
