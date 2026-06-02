package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := false

	for _, symbol := range s {
		if symbol == '-' && result == 0 {
			negative = true
		}
		if symbol >= '0' && symbol <= '9' {
			result = result*10 + int(symbol-'0')
		}
	}
	if negative {
		return -result
	}
	return result
}
