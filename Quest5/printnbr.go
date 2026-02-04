package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		var temp string
		for ; n != 0; n = n / 10 {
			temp = temp + string(rune(-(n%10)+'0'))
		}
		for i := len(temp) - 1; i >= 0; i-- {
			z01.PrintRune(rune(temp[i]))
		}
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		var temp string
		for ; n != 0; n = n / 10 {
			temp = temp + string(rune(n%10+'0'))
		}
		for i := len(temp) - 1; i >= 0; i-- {
			z01.PrintRune(rune(temp[i]))
		}
	}
}
