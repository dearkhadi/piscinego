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

	// Использование int64, чтобы избежать переполнения при -MinInt
	n := int64(nbr)

	// 2. Обработка знака
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// 3. Перевод системы счисления
	if n == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	var result []rune
	blen := int64(baseLen)

	for n > 0 {
		remainder := n % blen
		result = append(result, rune(base[remainder]))
		n /= blen
	}

	// 4. Вывод результата в обратном порядке
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

// Вспомогательная функция для вывода NV
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
