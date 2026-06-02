package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	}

	if nb > 2 {
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}
