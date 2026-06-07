package piscine

func RockAndRoll(n int) string {
	answer := ""

	if n < 0 {
		answer = "error: number is negative\n"
	} else if n%2 != 0 && n%3 != 0 {
		answer = "error: number is negative\n"
	}

	if n%2 == 0 && n%3 == 0 {
		answer = "rock and roll\n"
	} else if n%2 == 0 {
		answer = "rock\n"
	} else if n%3 == 0 {
		answer = "roll\n"
	}

	return answer
}
