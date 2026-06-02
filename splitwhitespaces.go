package piscine

func SplitWhiteSpaces(s string) []string {
	answerArray := []string{}
	var word string

	for _, letter := range s {
		if letter != ' ' && letter != '\t' && letter != '\n' {
			word += string(letter)
		} else if (letter == ' ' || letter == '\t' || letter == '\n') && word != "" {
			answerArray = append(answerArray, word)
			word = ""
		}
	}
	if word != "" {
		answerArray = append(answerArray, word)
	}
	return answerArray
}
