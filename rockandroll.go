package piscine

func RockAndRoll(n int) string {
	answer := ""

	if n < 0 {
		answer = "error: number is negative"
	} else if n%2 != 0 && n%3 != 0 {
		answer = "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {
		answer = "rock and roll"
	} else if n%2 == 0 {
		answer = "rock"
	} else if n%3 == 0 {
		answer = "roll"
	}

	return answer
}
