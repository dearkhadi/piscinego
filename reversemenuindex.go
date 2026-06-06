package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	answer := make([]string, n)
	for i := range menu {
		answer[n-1-i] = menu[i]
	}
	return answer
}
