package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	start := 0

	for i := 0; i <= len(s)-len(sep); {
		match := true

		for j := 0; j < len(sep); j++ {
			if s[i+j] != sep[j] {
				match = false
				break
			}
		}

		if match {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}

	result = append(result, s[start:])
	return result
}
