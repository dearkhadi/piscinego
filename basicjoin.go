package piscine

func BasicJoin(elems []string) string {
	result := ""

	for _, word := range elems {
		result += word
	}

	return result
}
