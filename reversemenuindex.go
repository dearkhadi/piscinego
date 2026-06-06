package piscine

func ReverseMenuIndex(menu []string) []string {
	answer := make([]string, len(menu))

	for i, str := range menu {
		answer[len(answer)-1-i] = str
	}

	return answer
}
