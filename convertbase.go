package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)

	if n == 0 {
		return string(baseTo[0])
	}

	baseLen := len(baseTo)
	result := ""

	for n > 0 {
		result = string(baseTo[n%baseLen]) + result
		n /= baseLen
	}

	return result
}
