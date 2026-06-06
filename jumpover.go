package piscine

func JumpOver(str string) string {
	myArr := []rune(str)
	answer := ""

	for i := 2; i < len(myArr); i += 3 {
		answer = answer + string(myArr[i])
	}

	answer = answer + "\n"
	return answer
}
