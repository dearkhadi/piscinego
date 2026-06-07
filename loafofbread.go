package piscine

func LoafOfBread(str string) string {
	charCount := 0
	for _, r := range str {
		if r != ' ' {
			charCount++
		}
	}

	if charCount < 5 {
		return "Invalid Output\n"
	}

	var res []rune
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			res = append(res, rune(str[i]))
			count++
		}

		if count == 5 {
			count = 0
			if i+1 < len(str) {
				i++
			}
			if i+1 < len(str) {
				res = append(res, ' ')
			}
		}
	}
	res = append(res, '\n')

	return string(res)
}
