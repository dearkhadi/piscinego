package piscine

func Abort(a, b, c, d, e int) int {
	answer := 0
	myArr := []int{a, b, c, d, e}

	for j := len(myArr) - 1; j > 0; j-- {
		sorted := false
		for i := len(myArr) - 1; i > len(myArr)-j-1; i-- {
			if myArr[i] < myArr[i-1] {
				myArr[i], myArr[i-1] = myArr[i-1], myArr[i]
				sorted = true
			}
		}
		if sorted == false {
			break
		}
	}

	answer = myArr[2]

	return answer
}
