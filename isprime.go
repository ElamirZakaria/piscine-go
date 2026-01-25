package piscine

func IsPrime(nb int) bool {
	result := true
	if nb == 2 {
		return true
	} else if nb < 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if i*i > nb {
			break
		}
		if nb%i == 0 {
			// fmt.Println(nb)
			// fmt.Println(i)
			result = false
			break
		} else {
			result = true
			break
		}

	}
	return result
}
