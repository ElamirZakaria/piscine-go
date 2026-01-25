package piscine

import "fmt"

func IsPrime(nb int) bool {
	result := true
	if nb == 2 || nb == 0 {
		return true
	} else if nb < 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			fmt.Println(i)
			fmt.Println(nb)
			result = false
			break
		} else {
			result = true
		}
	}
	return result
}
