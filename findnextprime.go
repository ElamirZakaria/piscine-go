package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	} else if nb%2 == 0 {
		nb++
	}
	for i := nb; i < 2147483647; i += 2 {
		if IsPrime(i) == true {
			return i
		} else {
		}
	}
	return nb
}
