package piscine

func Compare(a, b string) int {
	runeA := []rune(a)
	runeB := []rune(b)
	lenA := len(runeA)
	lenB := len(runeB)
	minLen := 0

	if lenA == lenB {
		minLen = lenA
	} else if lenA > lenB {
		minLen = lenB
	} else {
		minLen = lenA
	}

	for i := 0; i < minLen; i++ {
		if runeA[i] == runeB[i] {
		} else if runeA[i] < runeB[i] {
			return -1
		} else if runeA[i] > runeB[i] {
			return 1
		}
	}

	if lenA > lenB {
		return 1
	}

	return 0
}
