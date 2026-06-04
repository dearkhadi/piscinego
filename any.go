package piscine

func Any(f func(string) bool, a []string) bool {
	for _, str := range a {
		if f(str) == true {
			return true
		}
	}
	return false
}
