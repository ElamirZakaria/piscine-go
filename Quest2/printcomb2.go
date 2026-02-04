package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for a := '0'; a <= '9'; a++ {
					// if a < j {
					// 	continue
					// }
					// if k < i {
					// 	continue
					// }

					numberi := i - '0'
					numberj := j - '0'
					numberk := k - '0'
					numbera := a - '0'
					leftterm := numberi*10 + numberj
					rightterm := numberk*10 + numbera
					if leftterm >= rightterm {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(a)
					if leftterm == 98 && rightterm == 99 {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

	z01.PrintRune('\n')
}
