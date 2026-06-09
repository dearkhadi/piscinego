package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// 1. Валидация базы
	baseLen := len(base)
	if baseLen < 2 {
		printStr("NV")
		return
	}

	for i := 0; i < baseLen; i++ {
		if base[i] == '+' || base[i] == '-' {
			printStr("NV")
			return
		}
		for j := i + 1; j < baseLen; j++ {
			if base[i] == base[j] {
				printStr("NV")
				return
			}
		}
	}

	// 2. Обработка нуля
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// 3. Вывод знака для отрицательных чисел
	if nbr < 0 {
		z01.PrintRune('-')
	}

	var result []rune

	// 4. Разложение числа (работает одинаково хорошо и для положительных, и для MinInt)
	for nbr != 0 {
		remainder := nbr % baseLen
		// Если остаток отрицательный (для отрицательных nbr), делаем его положительным
		if remainder < 0 {
			remainder = -remainder
		}
		result = append(result, rune(base[remainder]))
		nbr /= baseLen
	}

	// 5. Вывод слайса в обратном порядке
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
