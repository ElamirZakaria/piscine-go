package piscine

func IsPrime(nb int) bool {
	result := true
	if nb == 2 {
		return true
	} else if nb < 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if i*i == nb {
			result = false
			break
		} else if i*i > nb {
			break
		}
	}
	return result
}
