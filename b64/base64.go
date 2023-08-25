package b64

import (
	"fmt"
	"strconv"
)

const ChunkSize = 6

func Encode(data string) (string, error) {
	var (
		bin       string
		result    string
		asciiData = []rune(data)
	)

	for _, char := range asciiData {
		bin += asciiToBin(char)
	}

	for _, chunk := range chunk(bin) {
		a, err := binToAscii(chunk)
		if err != nil {
			return "", err
		}
		r, err := getAlphaByIndex(int(a))
		if err != nil {
			return "", err
		}
		result += r
	}

	return padOutput(string(result)), nil
}

func Decode(data string) (string, error) {
	return "", nil
}

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

func chunk(str string) []string {
	var chunks []string
	str = padRight(str)

	for i := 0; i < len(str); i += ChunkSize {
		end := i + ChunkSize
		if end > len(str) {
			end = len(str)
		}
		chunks = append(chunks, str[i:end])
	}

	return chunks
}
