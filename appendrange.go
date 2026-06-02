package piscine

func AppendRange(min, max int) []int {
	myArray := []int{}

	if min >= max {
		return nil
	}

	for i := min; i < max; i++ {
		myArray = append(myArray, i)
	}
	return myArray
}
