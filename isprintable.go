package piscine

func IsPrintable(s string) bool {
	for _, symbol := range s {
		if symbol < ' ' || symbol > '~' {
			return false
		}
	}
	return true
}
