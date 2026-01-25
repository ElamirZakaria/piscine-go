package piscine

func IsPrime(nb int) bool {
	result := true
	if nb == 0 {
		return false
	} else if nb == 2 {
		return true
	} else if nb < 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {

			result = false
			break
		} else {
			result = true
		}
	}
	return result
}
