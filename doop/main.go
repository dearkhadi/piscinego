package main

import (
	"os"
)

func parseAndCheck(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0, false
	}

	var res uint64
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := uint64(s[i] - '0')
		next := res*10 + digit
		if next < res || (next-digit)/10 != res {
			return 0, false
		}
		res = next
	}

	if neg {
		if res > 9223372036854775808 {
			return 0, false
		}
		return -int64(res), true
	} else {
		if res > 9223372036854775807 {
			return 0, false
		}
		return int64(res), true
	}
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		if n == -9223372036854775808 {
			return "-9223372036854775808"
		}
		n = -n
	}

	var buf [20]byte
	i := len(buf) - 1
	for n > 0 {
		buf[i] = byte('0' + (n % 10))
		n /= 10
		i--
	}
	if neg {
		buf[i] = '-'
		i--
	}
	return string(buf[i+1:])
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	val1, ok1 := parseAndCheck(args[0])
	val2, ok2 := parseAndCheck(args[2])
	if !ok1 || !ok2 {
		return
	}

	op := args[1]

	switch op {
	case "+":
		if (val2 > 0 && val1 > 9223372036854775807-val2) || (val2 < 0 && val1 < -9223372036854775808-val2) {
			return
		}
		printStr(itoa(val1+val2) + "\n")

	case "-":
		if (val2 < 0 && val1 > 9223372036854775807+val2) || (val2 > 0 && val1 < -9223372036854775808+val2) {
			return
		}
		printStr(itoa(val1-val2) + "\n")

	case "*":
		if val1 != 0 && val2 != 0 {
			res := val1 * val2
			if res/val1 != val2 {
				return
			}
		}
		printStr(itoa(val1*val2) + "\n")

	case "/":
		if val2 == 0 {
			printStr("No division by 0\n")
			return
		}
		if val1 == -9223372036854775808 && val2 == -1 {
			return
		}
		printStr(itoa(val1/val2) + "\n")

	case "%":
		if val2 == 0 {
			printStr("No modulo by 0\n")
			return
		}
		if val1 == -9223372036854775808 && val2 == -1 {
			return
		}
		printStr(itoa(val1%val2) + "\n")

	default:
		return
	}
}
