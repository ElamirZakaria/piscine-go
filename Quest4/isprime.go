package piscine

func IsPrime(nb int) bool {
	result := true
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	} else {
		for i := 2; i <= nb; i++ {
			if nb%i == 0 && i != nb {

				result = false
				break
			} else {
				result = true
			}
		}
	}
	return result
}
