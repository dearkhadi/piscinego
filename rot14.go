func Rot14(s string) string {
	myArr := []rune(s)

	for index, rune := range myArr {
		if rune >= 'a' && rune <= 'z' {
			rune += 14
			if rune > 'z' {
				rune -= 26
			}
			myArr[index] = rune
		} else if rune >= 'A' && rune <= 'Z' {
			rune += 14
			if rune > 'Z' {
				rune -= 26
			}
			myArr[index] = rune
		}
	}

	return string(myArr)
}
