package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	value := 'F'
	if nb >= 0 {
		value = 'F'
	} else {
		value = 'T'
	}

	z01.PrintRune(value)
	z01.PrintRune('\n')
}
