package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var res string
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}

		res += string(str[i])
		count++

		if count == 5 {
			res += " "
			count = 0
			i++
		}
	}

	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res + "\n"
}
