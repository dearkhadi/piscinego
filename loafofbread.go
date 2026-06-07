package piscine

func LoafOfBread(str string) string {
	var res []rune
	charCount := 0
	for _, r := range str {
		if r != ' ' {
			charCount++
		}
	}

	if charCount < 5 {
		return "Invalid Output\n"
	}

	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			res = append(res, rune(str[i]))
			count++
		}

		if count == 5 {
			count = 0
			// Пропускаем 1 символ
			if i+1 < len(str) {
				i++
			}
			// Добавляем пробел только если впереди еще остались символы
			if i+1 < len(str) {
				res = append(res, ' ')
			}
		}
	}
	res = append(res, '\n')

	return string(res)
}
