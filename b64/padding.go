package b64

func padOutput(output string) string {
	var result string

	r := len(output) % 4

	if r > 0 {
		i := 0
		for i < 4-r {
			result += "="
			i++
		}
	}

	return output + result
}

func padLeft8bitBinary(bin string) string {
	var result string
	r := 8 - len(bin)

	for i := 0; i < r; i++ {
		result += "0"
	}

	return result + bin
}

func padRight(bin string) string {
	r := len(bin) % 6

	if r > 0 {
		i := 0
		for i < 6-r {
			bin += "0"
			i++
		}
	}

	return bin
}
