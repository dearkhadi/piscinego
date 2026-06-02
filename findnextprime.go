package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	if nb > 2 {
	cycle:
		for {
			for y := 2; y*y <= nb; y++ {
				if nb%y == 0 {
					nb++
					continue cycle
				}
			}
			return nb
		}
	}

	return nb
}
