package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)
	if len(runes) < 5 {
		return "Invalid Output\n"
	}

	var res string
	i := 0
	firstBlock := true

	for i < len(runes) {
		if !firstBlock {
			res += " "
		}

		count := 0
		for count < 5 && i < len(runes) {
			if runes[i] != ' ' {
				res += string(runes[i])
				count++
			}
			i++
		}

		firstBlock = false

		if i < len(runes) {
			i++
		}
	}
	return res + "\n"
}
