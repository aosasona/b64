package b64

func Encode(data string) (string, error) {
	var (
		bin       string
		result    string
		asciiData = []rune(data)
	)

	// convert each character to its 8-bit binary value and concatenate them to a single string
	for _, char := range asciiData {
		bin += asciiToBin(char)
	}

	for _, chunk := range chunk(bin) {
		// take the 6-bit binary chunk and convert it to the ascii equivalent
		a, err := binToAscii(chunk)
		if err != nil {
			return "", err
		}

		// take the ascii equivalent and convert it to the base64 alphabet equivalent
		r, err := getAlphaByIndex(int(a))
		if err != nil {
			return "", err
		}
		result += r
	}

	return padOutput(result), nil
}

func Decode(data string) (string, error) {
	var (
		bin    string
		result string
	)

	// convert each character to its 6-bit binary value and concatenate them
	for _, char := range data {
		// avoid padding
		if char == '=' {
			continue
		}

		bIdx, err := getIndexByAlpha(string(char))
		if err != nil {
			return "", err
		}

		// convert the index to 6-bit binary value and pad it to 6 bits
		bin += padLeft(intToBin(bIdx))
	}

	// convert the whole binary value to 8-bit binary chunks and convert each chunk to its ASCII equivalent
	for _, chunk := range chunk8bit(bin) {
		b, err := binToAscii(chunk)
		if err != nil {
			return "", err
		}

		// skip null bytes
		if b == 0 {
			continue
		}

		result += string(b)
	}

	return string(result), nil
}
