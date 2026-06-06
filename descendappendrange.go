package piscine

func DescendAppendRange(max, min int) []int {
	answer := []int{}

	if min >= max {
		return answer
	}

	for i := max; i > min; i-- {
		answer = append(answer, i)
	}

	return answer
}
