package piscine

func Capitalize(s string) string {
	myVar := []rune(s)

	for index, symbol := range myVar {

		isLower := symbol >= 'a' && symbol <= 'z'
		isUpper := symbol >= 'A' && symbol <= 'Z'

		if index == 0 {
			if isLower {
				myVar[index] = symbol - 32
			}
		} else {
			prev := myVar[index-1]

			isAlphaNum := (prev >= 'a' && prev <= 'z') || (prev >= 'A' && prev <= 'Z') || (prev >= '0' && prev <= '9')

			if !isAlphaNum && isLower {
				myVar[index] = symbol - 32
			}

			if isAlphaNum && isUpper {
				myVar[index] = symbol + 32
			}
		}
	}
	return string(myVar)
}
