package piscine

func Unmatch(a []int) int {
	myAnswer := -1
	lenArr := len(a)

	for i := 0; i < lenArr; i++ {
		checkNum := a[i]
		counter := 0

		for _, num := range a {
			if checkNum == num {
				counter++
			}
		}

		if counter%2 != 0 {
			myAnswer = checkNum
			break
		}
	}

	return myAnswer
}
