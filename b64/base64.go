package b64

import "strconv"

func Encode(data string) (string, error) {
	_ = []byte(data)
	return "", nil
}

func Decode(data string) (string, error) {
	return "", nil
}

func toCharCode(str rune) int {
	return int(str)
}

func toBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
