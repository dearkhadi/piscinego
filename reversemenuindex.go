package piscine

func ReverseMenuIndex(menu []string) []string {
	answer := make([]string, len(menu))
	n := len(menu) // Сохраним длину для удобства

	for i := 0; i < n; i++ {
		answer[n-1-i] = menu[i]
	}

	return answer
}
