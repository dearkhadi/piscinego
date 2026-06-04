package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0

	for _, str := range tab {
		if f(str) == true {
			result += 1
		}
	}
	return result
}
