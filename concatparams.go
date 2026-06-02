package piscine

func ConcatParams(args []string) string {
	var answer string

	for index, str := range args {
		if index >= 0 && index < len(args)-1 {
			answer = answer + string(str) + "\n"
		} else if index == len(args)-1 {
			answer = answer + string(str)
		}
	}
	return answer
}
