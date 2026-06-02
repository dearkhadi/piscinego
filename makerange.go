package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	myArray := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		myArray[i] = min + i
	}
	return myArray
}
