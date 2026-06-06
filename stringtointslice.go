package piscine

func StringToIntSlice(str string) []int {
	myVar := []int{}

	for _, s := range str {
		myVar = append(myVar, int(s))
	}
	return myVar
}
