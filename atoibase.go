package piscine

func AtoiBase(s string, base string) int {
	baseLen := len(base)

	if baseLen < 2 {
		return 0
	}

	for i := 0; i < baseLen; i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < baseLen; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	result := 0

	for i := 0; i < len(s); i++ {
		val := -1
		for index, char := range base {
			if rune(s[i]) == char {
				val = index
				break
			}
		}

		if val == -1 {
			return 0
		}

		result = result*baseLen + val
	}

	return result
}
