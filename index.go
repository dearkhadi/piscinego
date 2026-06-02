package piscine

func Index(s string, toFind string) int {
	findPlace := 0
	found := true
	myVar := []rune(s)
	myVar2 := []rune(toFind)

	for i := 0; i < len(myVar)-len(myVar2); i++ {
		for j := 0; j < len(myVar2); j++ {
			if myVar[i+j] == myVar2[j] {
				found = true
				findPlace = i
			} else {
				found = false
			}
		}
		if found == true {
			return findPlace
		}
	}
	return -1
}
