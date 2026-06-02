package piscine

func Swap(a *int, b *int) {
	i := *a
	y := *b
	*a = y
	*b = i
}
