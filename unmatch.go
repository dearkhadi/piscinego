package piscine

func Unmatch(a []int) int {
	myAnswer := -1
	lenArr := len(a)

	for i := 0; i < lenArr; i++ {
		checkNum := a[i]

		for _, num := range a {
			if !(checkNum == num) {
				myAnswer = checkNum
				break
			}
		}
	}

	return myAnswer
}
