package piscine

func FindNextPrime(nb int) int {
	if nb%2 == 0 {
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
