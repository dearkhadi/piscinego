package piscine

func Join(strs []string, sep string) string {
	result := ""

	for index, word := range strs {
		if index > 0 {
			result += sep
		}

		result += word
	}

	return result
}
