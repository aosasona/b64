package b64

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
	var result, bin string

	for _, char := range data {
		if char == '=' {
			continue
		}

		bIdx, err := getIndexByAlpha(string(char))
		if err != nil {
			return "", err
		}

		bin += padLeft(intToBin(bIdx))
	}

	for _, chunk := range chunk8bit(bin) {
		b, err := binToAscii(chunk)
		if err != nil {
			return "", err
		}
		result += string(b)
	}

	return result, nil
}
