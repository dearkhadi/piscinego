package piscine

func StringToIntSlice(str string) []int {
	var myVar []int

	for _, s := range str {
		myVar = append(myVar, int(s))
	}
	if len(myVar) > 0 {
		return myVar
	}
	return nil
}
