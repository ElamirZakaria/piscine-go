package piscine

func FindNextPrime(nb int) int {
	for {
		if IsPrime(nb) == true {
			return nb
		} else {
			nb++
		}
	}
}
