package piscine

func StrRev(s string) string {
	a := []rune(s)
	b := len(a)
	i := 0

	for i < b/2 {
		c := a[i]
		a[i] = a[len(a)-1-i]
		a[len(a)-1-i] = c
		i = i + 1
	}
	return string(a)
}
