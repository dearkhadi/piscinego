package piscine

func LoafOfBread(str string) string {
	count := 0
	for _, ch := range str {
		if ch != ' ' {
			count++
		}
	}
	if count == 0 {
		return "\n"
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0
	firstWord := true

	for i < len(str) {
		word := ""
		for i < len(str) && len(word) < 5 {
			if str[i] != ' ' {
				word += string(str[i])
			}
			i++
		}
		if word == "" {
			break
		}
		if !firstWord {
			result += " "
		}
		result += word
		firstWord = false
		if len(word) == 5 && i < len(str) {
			i++
		}
	}

	return result + "\n"
}
