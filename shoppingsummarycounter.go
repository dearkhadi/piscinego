package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	myStrArr := []string{}
	answer := make(map[string]int)

	word := ""
	for _, rune := range str {
		if rune != ' ' {
			word = word + string(rune)
		} else {
			myStrArr = append(myStrArr, word)
			word = ""
		}
	}

	myStrArr = append(myStrArr, word)

	for _, word := range myStrArr {
		answer[word] = answer[word] + 1
	}

	return answer
}
