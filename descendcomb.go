package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	a := [100]int{}
	j := 0

	for i := 99; i >= 0; i-- {
		a[j] = i
		j++
	}
	for i := 0; i < 99; i++ {
		for j = i + 1; j <= 99; j++ {
			if a[i] < 10 {
				z01.PrintRune('0')
				z01.PrintRune(rune(a[i]) + '0')
				z01.PrintRune(rune(' '))
			} else {
				z01.PrintRune(rune(a[i]/10) + '0')
				z01.PrintRune(rune(a[i]%10) + '0')
				z01.PrintRune(rune(' '))

			}
			if a[j] < 10 {
				z01.PrintRune('0')
				z01.PrintRune(rune(a[j]) + '0')
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			} else {
				z01.PrintRune(rune(a[j]/10) + '0')
				z01.PrintRune(rune(a[j]%10) + '0')
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}

		}
	}
}
