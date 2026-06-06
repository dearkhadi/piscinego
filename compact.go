package piscine

func Compact(ptr *[]string) int {
	for i := 0; i < len((*ptr)); i++ {
		if (*ptr)[i] == "" {
			*ptr = append((*ptr)[:i], (*ptr)[i+1:]...)
			i--
		}
	}

	return len((*ptr))
}
